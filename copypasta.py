import subprocess
import sys
import time
from colorama import Fore, Style
from halo import Halo  # For spinner

def run_wget(command):
    print(Fore.GREEN + f"Running command: {command}" + Style.RESET_ALL)
    spinner = Halo(text='Downloading...', spinner='dots')
    spinner.start()
    
    process = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    
    for line in process.stdout:
        formatted_line = format_output(line.decode().strip())
        print(formatted_line)
    
    process.wait()
    spinner.stop()
    
    if process.returncode != 0:
        print(Fore.RED + "Error occurred!" + Style.RESET_ALL)

def format_output(line):
    if line.startswith("["):
        return Fore.GREEN + line + Style.RESET_ALL
    elif "Resolving" in line:
        return Fore.CYAN + line + Style.RESET_ALL  # Resolving IP addresses
    elif "Connecting" in line:
        return Fore.YELLOW + line + Style.RESET_ALL  # Connection info
    elif "HTTP request sent" in line:
        return Fore.MAGENTA + line + Style.RESET_ALL  # HTTP request status
    elif "Not Modified" in line:
        return Fore.BLUE + line + Style.RESET_ALL  # Not modified response
    else:
        return Fore.WHITE + line + Style.RESET_ALL  # Default color for other lines

def main():
    if len(sys.argv) < 3:
        print(Fore.YELLOW + "Usage: copypasta <url> <type>" + Style.RESET_ALL)
        return

    url = sys.argv[1]
    command_type = sys.argv[2]

    if command_type == "mirror":
        command = f"wget --mirror --convert-links --adjust-extension {url}"
    elif command_type == "page":
        command = f"wget --mirror --convert-links --adjust-extension --page-requisites {url}"
    else:
        print(Fore.RED + "Invalid command type!" + Style.RESET_ALL)
        return

    run_wget(command)

if __name__ == "__main__":
    main()
