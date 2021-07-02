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


# Remove logrus add-on if not seleted
if '{{ cookiecutter.use_logrus }}'.lower() == 'no':
    print_log('Removing logrus add-on')
    remove_files("helpers/log")

# Download the missing dependencies
print_log('Downloading dependencies')
os.system("go mod tidy")

# Format code
print_log('Formating code')
os.system("go fmt ./...")

# # Use this snippet to remove files/folders based on
# # cookiecutter.variables (cf pre_gen_project.py hook)
# if '{{ cookiecutter.web_variant }}' == 'n':
#     print_log('Removing web variant related files')
#     remove_files('/path/to/remove')
