import os
import shutil

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)


def remove_logrus_files():
    """
    Removes log helper folder
    """
    shutil.rmtree(os.path.join(
        PROJECT_DIRECTORY, "helpers/log"
    ))


# Remove logrus utils if not seleted
if '{{ cookiecutter.use_logrus }}'.lower() != 'y':
    remove_logrus_files()

# Download the missing dependencies
os.system("go mod tidy")

# Format code
os.system("go fmt ./...")
