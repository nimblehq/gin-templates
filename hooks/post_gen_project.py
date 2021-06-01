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


def print_log(message):
    """
    Print log with color
    """
    CYELLOW = '\33[33m'  # YELLOW color
    CEND = '\033[0m'  # END color
    print(CYELLOW + message + CEND)


# Remove logrus add-on if not seleted
if '{{ cookiecutter.use_logrus }}'.lower() != 'y':
    print_log('Removing logrus add-on')
    remove_logrus_files()

# Download the missing dependencies
print_log('Downloading dependencies')
os.system("go mod tidy")

# Format code
print_log('Formating code')
os.system("go fmt ./...")
