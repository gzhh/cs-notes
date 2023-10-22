
"""
The import system
- https://docs.python.org/3/reference/import.html

importlib — The implementation of import
- https://docs.python.org/3/library/importlib.html
"""

print('test.py name: ', __name__)

if __name__ == '__main__':
    from task.task1 import task_task1_func
    task_task1_func()

    from task import task1 
    task1.task_task1_func()

