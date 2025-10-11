<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [pyenv (Python version management)](#pyenv-python-version-management)
  - [Usage](#usage)
- [pip (Python package management)](#pip-python-package-management)
  - [Installing specific package version](#installing-specific-package-version)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Python
## pyenv (Python version management)
- https://github.com/pyenv/pyenv
- see also
  - https://github.com/pyenv/pyenv-virtualenv
  - https://docs.python.org/3/library/venv.html

### Usage
List all Python versions available to pyenv
- `pyenv versions`

Install a Python version
- `pyenv install 3.10.4`

Set the global Python version
- `pyenv global 3.10.4`

### pyenv install pip
确保 pyenv 和对应的 Python 版本安装正常，然后使用 ensurepip 和 pip 命令来升级 pip。

```
pyenv install <version>  # 安装特定的 Python 版本
pyenv global <version>   # 或者 pyenv local <version>

python -m ensurepip --upgrade
python -m pip install --upgrade pip
```

运行以下命令以确保 pyenv shims 目录正确更新。

`pyenv rehash`


## pip (Python package management)
- https://en.wikipedia.org/wiki/Pip_(package_manager)
- https://github.com/pypa/pip
- https://pip.pypa.io/en
- see also
  - https://github.com/pypa/pipenv
  - https://github.com/pdm-project/pdm
  - https://github.com/python-poetry/poetry

### Installing specific package version
Example
- `pip install --force-reinstall -v "MySQL_python==1.2.2"`

Ref
- https://stackoverflow.com/questions/5226311/installing-specific-package-version-with-pip


## uv (Python version and package management)
- https://github.com/astral-sh/uv

Best Practice
- https://docs.astral.sh/uv/guides/scripts/
- uv ：新一代的 Python 包管理工具 https://liujiacai.net/blog/2025/08/03/uv-intro/
- https://www.bilibili.com/video/BV1Stwfe1E7s


## gunicorn
- gunicorn https://github.com/benoitc/gunicorn
