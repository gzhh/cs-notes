## pip 使用
pip 升级
- pip install --upgrade pip
- pip install -U pip

pip 安装具体包
- pip install sampleproject
- python -m pip install sampleproject
- python -m pip install git+https://github.com/pypa/sampleproject.git@main

pip 根据文件安装包
- pip install -r requirements.txt

命令freeze 让pip将项目中当前安装的所有包的名称都写入文件 requirements.txt
- pip freeze > requirements.txt

## venv 和 virtualenv 使用
- python -m venv /path/to/new/virtual/environment
- python -m virtualenv /path/to/new/virtual/environment
