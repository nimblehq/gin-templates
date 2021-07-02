import os
import shutil

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

def remove_files(path):
    """
    Removes files in path
    """
    shutil.rmtree(os.path.join(
        PROJECT_DIRECTORY, path
    ))

# Print log with color
def print_log(message):
    """
    Print log with color
    """
    CYELLOW = '\33[33m'  # YELLOW color
    CEND = '\033[0m'  # END color
    print(CYELLOW + message + CEND)


def remove_heroku_files():
    """
    Removes heroku related files
    """
    shutil.rmtree(os.path.join(
        PROJECT_DIRECTORY, "deploy/heroku"
    ))
    os.remove(os.path.join(
        PROJECT_DIRECTORY, ".github/workflows/deploy.yml"
    ))


def remove_empty_folders():
    """
    List all empty folders and remove them
    """
    for root, dirs, files in os.walk(PROJECT_DIRECTORY, topdown=False):
        for dir in dirs:
            if os.path.isdir(os.path.join(root, dir)):
                if (len(os.listdir(os.path.join(root, dir))) == 0):
                    print_log(f'Removing {dir} folder')
                    os.rmdir(os.path.join(root, dir))


# Remove logrus add-on if not seleted
if '{{ cookiecutter.use_logrus }}'.lower() == 'no':
    print_log('Removing logrus add-on')
    remove_files("helpers/log")

# Remove heroku add-on if not seleted
if '{{ cookiecutter.use_heroku }}'.lower() == 'no':
    print_log('Removing heroku add-on')
    remove_heroku_files()

# Download the missing dependencies
print_log('Downloading dependencies')
os.system("go mod tidy")

# Format code
print_log('Formating code')
os.system("go fmt ./...")

# Remove empty folders
remove_empty_folders()
