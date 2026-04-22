import uuid
from enum import Enum
from typing import List, Optional

class TaskStatus(Enum):
    BLOCKED = "BLOCKED"
    PENDING = "PENDING"
    PROCESSING = "PROCESSING"
    COMPLETED = "COMPLETED"
    FAILED = "FAILED"
    CANCELLED = "CANCELLED"

class Task:
    def __init__(self, id: str, name: str, status: TaskStatus = TaskStatus.PENDING):
        self.id = id
        self.name = name
        self.status = status
        self.dependencies: List[str] = []

class DAGManager:
    """
    TaskMaster DAG Logic for AGENCIA01
    Manages task dependencies and unlocking cascades.
    """
    def __init__(self):
        self.tasks: dict[str, Task] = {}

    def add_task(self, name: str, dependencies: List[str] = None) -> str:
        task_id = str(uuid.uuid4())[:8]
        status = TaskStatus.BLOCKED if dependencies else TaskStatus.PENDING
        task = Task(task_id, name, status)
        if dependencies:
            task.dependencies = dependencies
        self.tasks[task_id] = task
        return task_id

    def complete_task(self, task_id: str):
        if task_id not in self.tasks:
            return
        
        self.tasks[task_id].status = TaskStatus.COMPLETED
        
        # Unlocking logic
        for tid, task in self.tasks.items():
            if task.status == TaskStatus.BLOCKED:
                if all(self.tasks[dep].status == TaskStatus.COMPLETED for dep in task.dependencies if dep in self.tasks):
                    task.status = TaskStatus.PENDING
                    print(f"[DAG] Task {tid} ({task.name}) UNLOCKED")

    def get_pending_tasks(self) -> List[Task]:
        return [t for t in self.tasks.values() if t.status == TaskStatus.PENDING]

if __name__ == "__main__":
    dag = DAGManager()
    t1 = dag.add_task("Setup Env")
    t2 = dag.add_task("Install Deps", dependencies=[t1])
    t3 = dag.add_task("Run App", dependencies=[t2])

    print(f"Initial Pending: {[t.name for t in dag.get_pending_tasks()]}")
    dag.complete_task(t1)
    print(f"After T1 Pending: {[t.name for t in dag.get_pending_tasks()]}")
    dag.complete_task(t2)
    print(f"After T2 Pending: {[t.name for t in dag.get_pending_tasks()]}")
