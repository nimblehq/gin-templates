import os
import shutil

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

# Removes log helper folder
def remove_logrus_files():
    shutil.rmtree(os.path.join(
        PROJECT_DIRECTORY, "helpers/log"
    ))

# Print log with color
def print_log(message):
    CYELLOW = '\33[33m'  # YELLOW color
    CEND = '\033[0m'  # END color
    print(CYELLOW + message + CEND)


# Remove logrus add-on if not seleted
if '{{ cookiecutter.use_logrus }}'.lower() == 'no':
    print_log('Removing logrus add-on')
    remove_logrus_files()

# Download the missing dependencies
print_log('Downloading dependencies')
os.system("go mod tidy")

# Format code
print_log('Formating code')
os.system("go fmt ./...")
