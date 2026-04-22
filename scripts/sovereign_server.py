from http.server import SimpleHTTPRequestHandler, HTTPServer
import os
import json
from urllib.parse import urlparse, parse_qs

PORT = 8000
DIRECTORY = "dashboard"
SQUAD_PATH = "squads/agencia-squad/prompts"

class SovereignHandler(SimpleHTTPRequestHandler):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, directory=DIRECTORY, **kwargs)

    def do_GET(self):
        parsed_path = urlparse(self.path)
        if parsed_path.path == '/inspect':
            params = parse_qs(parsed_path.query)
            agent_name = params.get('agent', [None])[0]
            
            if not agent_name:
                self.send_error(400, "Missin agent name")
                return

            # OpSec: Sanitização de path
            agent_name = os.path.basename(agent_name)
            prompt_file = os.path.join(SQUAD_PATH, f"{agent_name}.md")
            
            if os.path.exists(prompt_file):
                with open(prompt_file, 'r') as f:
                    content = f.read()
                
                self.send_response(200)
                self.send_header('Content-type', 'text/markdown')
                self.send_header('Access-Control-Allow-Origin', '*')
                self.end_headers()
                self.wfile.write(content.encode())
            else:
                self.send_error(404, "Agent prompt not found")
        else:
            super().do_GET()

    def do_POST(self):
        content_length = int(self.headers['Content-Length'])
        post_data = self.rfile.read(content_length)
        
        try:
            command = json.loads(post_data)
            cmd_type = command.get("command")
            cmd_dir = os.path.join(DIRECTORY, "data/commands")
            os.makedirs(cmd_dir, exist_ok=True)
            
            cmd_file = os.path.join(cmd_dir, f"last_cmd.json")
            with open(cmd_file, "w") as f:
                json.dump(command, f)
            
            self.send_response(200)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            self.wfile.write(json.dumps({"status": "success", "msg": f"Command {cmd_type} queued"}).encode())
        except Exception as e:
            self.send_response(500)
            self.end_headers()
            self.wfile.write(str(e).encode())

def run():
    server_address = ('', PORT)
    httpd = HTTPServer(server_address, SovereignHandler)
    print(f"[*] Sovereign Server V2 (Deep Inspection) rodando na porta {PORT}...")
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        httpd.server_close()

if __name__ == "__main__":
    run()
