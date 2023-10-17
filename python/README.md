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
  - https://twitter.com/river_leaves/status/1710489478924783778
  - https://github.com/pypa/pipenv
  - https://github.com/python-poetry/poetry
  - https://github.com/pdm-project/pdm

### Usage
List all Python versions available to pyenv
- `pyenv versions`

Install a Python version
- `pyenv install 3.10.4`

Set the global Python version
- `pyenv global 3.10.4`


## pip (Python package management)
- https://en.wikipedia.org/wiki/Pip_(package_manager)
- https://github.com/pypa/pip
- https://pip.pypa.io/en

### Installing specific package version
Example
- `pip install --force-reinstall -v "MySQL_python==1.2.2"`

Ref
- https://stackoverflow.com/questions/5226311/installing-specific-package-version-with-pip

