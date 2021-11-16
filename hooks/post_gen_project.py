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

def remove_file(path):
    """
    Removes file with path
    """
    os.remove(os.path.join(
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

    remove_files("deploy/heroku")
    remove_file(".github/workflows/deploy.yml")

# Remove bootstrap add-on if not seleted
if '{{ cookiecutter._bootstrap_addon }}' == 'no':
    print_log('Removing bootstrap add-on')
    remove_files("lib/web/assets/stylesheets/vendors/bootstrap")

# Remove tailwind add-on if not seleted
if '{{ cookiecutter._tailwind_addon }}' == 'no':
    print_log('Removing tailwind add-on')
    remove_files("lib/web/assets/stylesheets/vendors/tailwind")
    remove_file("tailwind.config.js")

# Remove github repository add-on if not seleted
if '{{ cookiecutter._github_addon }}' == 'no':
    print_log('Removing github repository add-on')
    remove_files(".github")

# Remove gitlab repository add-on if not seleted
if '{{ cookiecutter._gitlab_addon }}' == 'no':
    print_log('Removing gitlab repository add-on')
    remove_files(".gitlab")

# Remove .gitignore if not seleted any repository options
if '{{ cookiecutter._gitignore_addon }}' == 'no':
    print_log('Removing .gitignore')
    remove_file("tailwind.config.js")

# Remove web variant if not selected
if '{{ cookiecutter._web_variant }}' == 'no':
    print_log('Removing web variant related files')

    # Web folder
    remove_files('lib/web')
    
    # Config files
    remove_file('.eslintrc.json')
    remove_file('.npmrc')
    remove_file('package.json')
    remove_file('snowpack.config.js')
    remove_file('postcss.config.js')
    remove_file('tsconfig.json')

# Download the missing dependencies
print_log('Downloading dependencies')
os.system("go mod tidy")

# Format code
print_log('Formating code')
os.system("go fmt ./...")

# Remove empty folders
remove_empty_folders()
