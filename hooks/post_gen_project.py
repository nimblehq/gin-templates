import os
import sys

REMOVE_PATHS = [
    '{% if cookiecutter.variant == "Web" %} api.txt {% endif %}',
    '{% if cookiecutter.variant == "API" %} web.txt {% endif %}',
]

for path in REMOVE_PATHS:
    path = path.strip()
    if path and os.path.exists(path):
        if os.path.isdir(path):
            os.rmdir(path)
        else:
            os.unlink(path)