import os
import json
import time
import socket
import re
import sys
from datetime import datetime

base_path = "/home/kali/Área de trabalho/AGENCIA01"
if base_path not in sys.path:
    sys.path.append(base_path)
output_path = os.path.join(base_path, "dashboard/data/status.json")
missions_path = os.path.join(base_path, "dashboard/data/missions.json")
squad_yml_path = os.path.join(base_path, "squads/agencia-squad/squad.yml")
cmd_file_path = os.path.join(base_path, "dashboard/data/commands/last_cmd.json")

PROJ_ID = "AG_01_SQUAD"
FORBIDDEN_WORDS = [os.getlogin(), "/home/"]

def mask_content(data):
    str_data = json.dumps(data)
    for word in FORBIDDEN_WORDS:
        str_data = str_data.replace(word, "***")
    return json.loads(str_data)

def check_tor():
    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.settimeout(0.5)
        s.connect(("127.0.0.1", 9050))
        s.close()
        return True
    except:
        return False

def get_agents():
    agents = []
    if not os.path.exists(squad_yml_path): return agents
    with open(squad_yml_path, 'r') as f:
        content = f.read()
        matches = re.finditer(r'- name: (.*?)\n\s+role: (.*?)\n', content)
        for m in matches:
            agents.append({
                "name": m.group(1).strip(),
                "role": m.group(2).strip(),
                "status": "idle",
                "prompt_ref": f"squads/agencia-squad/prompts/{m.group(1).strip()}.md"
            })
    return agents

from utils.dag import DAGManager, TaskStatus

def get_mission_data():
    if os.path.exists(missions_path):
        with open(missions_path, 'r') as f:
            missions = json.load(f)
            active = next((m for m in missions if m['status'] == "ACTIVE"), missions[-1])
            
            # Integrando Lógica de DAG (TaskMaster)
            steps = active.get('steps', [])
            dag = DAGManager()
            # Mapeia steps do JSON para o DAG
            for s in steps:
                t = dag.add_task(s['name'], s.get('depends_on', []))
                dag.tasks[s['id']] = dag.tasks.pop(t) # Mantém IDs originais
                dag.tasks[s['id']].id = s['id']
                if s['status'] == "done":
                    dag.tasks[s['id']].status = TaskStatus.COMPLETED
                elif s['status'] == "doing":
                    dag.tasks[s['id']].status = TaskStatus.PROCESSING

            doing_task = next((s for s in steps if s['status'] == "doing"), {"id": "IDLE", "name": "Waiting"})
            return active, doing_task
    return {}, {}

def process_commands():
    if os.path.exists(cmd_file_path):
        # ... logic for commands ...
        try:
            with open(cmd_file_path, 'r') as f: cmd = json.load(f)
            action = cmd.get("command")
            if action == "EMERGENCY_STOP":
                os.remove(cmd_file_path)
                sys.exit(0)
            os.remove(cmd_file_path)
        except: pass

def get_stats():
    try:
        import psutil
        cpu = psutil.cpu_percent()
        ram = psutil.virtual_memory().percent
    except: cpu, ram = 0.0, 0.0
    
    active_mission, active_task = get_mission_data()
    agents = get_agents()
    
    # Simulação de atividade por tarefa
    for agent in agents:
        if agent['name'] in ['coder_hacker', 'monitor'] and active_task['id'] != "IDLE":
            agent['status'] = 'active'
            
    return {
        "timestamp": datetime.now().strftime("%Y-%m-%dT%H:%M:%S.%f")[:-3] + "Z",
        "project_id": PROJ_ID,
        "mission_id": active_mission.get('id', 'N/A'),
        "active_task": active_task,
        "system": {
            "status": "online",
            "cpu": cpu,
            "ram": ram,
            "tor": check_tor()
        },
        "agents": agents,
        "missions": get_missions_raw()
    }

def get_missions_raw():
    if os.path.exists(missions_path):
        with open(missions_path, 'r') as f: return json.load(f)
    return []

def main():
    os.makedirs(os.path.dirname(output_path), exist_ok=True)
    os.makedirs(os.path.dirname(cmd_file_path), exist_ok=True)
    while True:
        try:
            process_commands()
            stats = get_stats()
            safe_stats = mask_content(stats)
            with open(output_path, "w") as f:
                json.dump(safe_stats, f, indent=2)
            time.sleep(1) # Maior precisão: refresh 1s
        except SystemExit: break
        except Exception as e: time.sleep(5)

if __name__ == "__main__":
    main()
