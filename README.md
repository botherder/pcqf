# pcqf

[![Go Report Card](https://goreportcard.com/badge/github.com/botherder/pcqf)](https://goreportcard.com/report/github.com/botherder/pcqf)

pcqf (PC Quick Forensics) is a simple tool to automate the acquisition of some evidence of compromise from Windows computers. pcqf is normally intended for trainers, researchers, and incident responders without a particular background in information security and computer forensics.

Often, it is not possible (because of logistical reasons, lack of appropriate hardware, or simply privacy issues) to obtain a full disk image of the computer. pcqf allows to gather sufficient data to initiate and investigation, while minimizing exposure of personal data and without requiring a particular expertise in computer forensics.

pcqf doesn't require any configuration or parameters, it just needs to be executed with Administrator privileges. Once launched, the software automatically harvests and collects copies of the executables of running processes and of the applications automatically starting at launch. Optionally, it can also take a full-memory dump.

[Download pcqf](https://github.com/botherder/pcqf/releases/latest)

## How to use

1. Download pcqf on a USB device. Make sure that the device has enough space to store all the acquisitions you are going to make. It is advisable to format the USB device as NTFS, in case you will end up dumping memory of computers with significant RAM.

2. Mount the USB device on the computer to inspect. Browse to the pcqf folder and double-click on the tool. It should ask you to allow the application to run with Administrator privileges, which are required to obtain a memory snapshot. On Mac computers, you will need to launch pcqf from the terminal with the commands `chmod +x pcqf` and `sudo ./pcqf`.

3. Wait for the tool to complete its execution. You will see some log messages displayed in console. Pay particular attention in case it mentions problems for example in relation to the generation of the memory dump.

4. Once completed, you will find a new folder called "acquisitions". Inside this folder you will see a folder for each acquisition you made. The folders will be named in the format `YYYY-MM-DD_\<COMPUTER NAME\>`. You can perform multiple acquisitions from the same computer: new folders will be distinguished by a numeric suffix.

5. Each acquisition folder will contain the following files:

    - A `profile.json` file containing basic information on the computer system.
    - A `process_list.json` file containing a list of running processes.
    - A `autoruns.json` file containing a list of all items with persistence on the system.
    - A `autoruns_bins/` folder containing copies of the files and executables marked for persistence in the previous JSON file.
    - A `process_bins/` folder containing copies of the executables of running processes.
    - If successful, a `memory/` folder will contain a physical memory dump as well as some metadata.

## Encryption & Potential Threats

Carrying the pcqf acquisitions on an unencrypted drive might expose yourself, and even more so those you acquired data from, to significant risk. For example, you might be stopped at a problematic border and your pcqf drive could be seized. The raw data might not only expose the purpose of your trip, but it will also likely contain very sensitive data (in the memory image, for example, one could find usernames & passwords, browsing history, and more).

Ideally you should have the drive fully encrypted, but that might not always be possible. You could also consider placing pcqf inside a [VeraCrypt](https://www.veracrypt.fr/) container and carry with it a copy of VeraCrypt to mount it. However, VeraCrypt containers are typically protected only by a password, which you might be forced to provide.

Alternatively, pcqf allows to encrypt each acquisition with a provided PGP public key. Preferably, this public key belongs to a keypair for which you don't possess, or at least carry, the private key. In this way, you would not be capable of decrypting the acquisitions data even under duress.

If you place a file called `public.asc` in the same folder as the pcqf executable, pcqf will automatically attempt to compress and encrypt each acquisition and delete the original unencrypted copies. The encrypted file will be named with an undescriptive unique identifier.

Bear in mind, it is always possible that at least some portion of the unencrypted data could be recovered through advanced forensics techniques - although we're working to mitigate that.

## Known issues

The memory acquisition does not work on Windows XP.
