# pcqf

[![Go Report Card](https://goreportcard.com/badge/github.com/botherder/pcqf)](https://goreportcard.com/report/github.com/botherder/pcqf)

pcqf (PC Quick Forensics) is a simple tool to automate the acquisition of some evidence of compromise from Windows computers. pcqf is normally intended for trainers, researchers, and incident responders without a particular background in information security and computer forensics.

Often, it is not possible (because of logistical reasons, lack of appropriate hardware, or simply privacy issues) to obtain a full disk image of the computer. pcqf allows to gather sufficient data to initiate and investigation, while minimizing exposure of personal data and without requiring a particular expertise in computer forensics.

pcqf doesn't require any configuration or parameters, it just needs to be executed with Administrator privileges. Once launched, the software automatically harvests and collects copies of the executables of running processes and of the applications automatically starting at launch. Optionally, it can also take a full-memory dump.

[Download pcqf](https://github.com/botherder/pcqf/releases/latest)

## Build

Executable binaries for Linux, Windows and Mac should be available in the [latest release](https://github.com/botherder/pcqf/releases/latest). In case you have issues running the binary you might want to build it by yourself.

In order to build pcqf you will need Go 1.15+ installed. You will also need to install `make`. When ready you can clone the repository and run any of the following commands, for your platform of choice:

    make linux
    make darwin
    make windows

These commands will generate binaries in a *build/* folder.

## How to use

1. Download pcqf on a USB device. Make sure that the device has enough space to store all the acquisitions you are going to make. It is advisable to format the USB device as NTFS, in case you will end up dumping memory of computers with significant RAM.

2. Mount the USB device on the computer to inspect. Browse to the pcqf folder and double-click on the tool. It should ask you to allow the application to run with Administrator privileges, which are required to obtain a memory snapshot. On Mac computers, you will need to launch pcqf from the terminal with the commands `chmod +x pcqf` and `sudo ./pcqf`.

3. Wait for the tool to complete its execution. You will see some log messages displayed in console. Pay particular attention in case it mentions problems for example in relation to the generation of the memory dump.

4. Once completed, you will find a new folder named with a UUID. You can perform multiple acquisitions from the same computer: each acquisition will have a different UUID.

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

Alternatively, pcqf allows to encrypt each acquisition with a provided [age](https://age-encryption.org) public key. Preferably, this public key belongs to a keypair for which the end-user does not possess, or at least carry, the private key. In this way, the end-user would not be able to decrypt the acquired data even under duress.

If you place a file called `key.txt` in the same folder as the pcqf executable, pcqf will automatically attempt to compress and encrypt each acquisition and delete the original unencrypted copies.

Once you have retrieved an encrypted acquisition file, you can decrypt it with age like so:

```
$ age --decrypt -i ~/path/to/privatekey.txt -o b51e49b9-02ae-457f-95d1-f83ae8029488.zip b51e49b9-02ae-457f-95d1-f83ae8029488.zip.age
```

Bear in mind, it is always possible that at least some portion of the unencrypted data could be recovered through advanced forensics techniques - although we're working to mitigate that.

## Known issues

The memory acquisition does not work on Windows XP.
