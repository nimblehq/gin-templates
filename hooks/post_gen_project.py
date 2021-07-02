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


def print_log(message):
    """
    Print log with color
    """
    CYELLOW = '\33[33m'  # YELLOW color
    CEND = '\033[0m'  # END color
    print(CYELLOW + message + CEND)


# # Use this snippet to remove files/folders based on
# # cookiecutter.variables (cf pre_gen_project.py hook)
# if '{{ cookiecutter.web_variant }}' == 'n':
#     print_log('Removing web variant related files')
#     remove_files('/path/to/remove')
