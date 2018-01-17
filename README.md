![Snoopdigg](https://github.com/botherder/snoopdigg/raw/master/graphics/icon%40128.png)

# Snoopdigg

Snoopdigg is a simple tool to automate some basic steps to acquire some
evidence of compromise from Windows computers. Snoopdigg is normally intended for
trainers, researchers, and incident responders without a particular background
in information security and computer forensics.

Snoopdigg doesn't require any configuration or parameters, it just needs to
be executed with Administrator privileges. Once launched, the software
automatically harvests and collects copies of the Windows executables that
maintain persistence on the system, and afterwards attempts at taking a full
snapshot of the memory.

Often, it is not possible (because of logistical reasons, lack of appropriate
hardware, or simply privacy issues) to do a full disk image of the computer.
Snoopdigg allows to at least fetch sufficient data to initiate an
investigation minimizing the exposure of personal information as well as
avoiding the need for the person performing the acquisition to be specifically
trained in using rather unfriendly tools.

## How to use

1. Extract this folder on a USB device. Make sure that the device has enough
space to store all the acquisitions you are going to make. It is advisable to
format the USB device as NTFS, in case you will end up dumping memory of
computers with significant RAM.

2. Mount the USB device on the computer to inspect. Browse to the Snoopdigg
folder and double-click on the tool. It should ask you to allow the application
to run with Administrator privileges, which are required to obtain a memory
snapshot.

3. Wait for the tool to complete its execution. You will see some log messages
displayed in console. Pay particular attention in case it mentions problems
for example in relation to the generation of the memory dump.

4. Once completed, you will find a new folder called "acquisitions". Inside this
folder you will see a folder for each acquisition you made. The folders will
be named in the format "YYYY-MM-DD_\<COMPUTER NAME\>". You can perform
multiple acquisitions from the same computer, new folders will be distinguished
by a numeric suffix.

5. Each acquisition folder will contain the following files:

    - A *profile.json* file containing basic information on the computer system.
    - A *processlist.json* file containing a list of running processes.
    - An *autoruns.json* file containing a list of all items with persistence on
      the system.
    - An *autoruns/* folder containing copies of the files and executables
      marked for persistence in the previous Json file.
    - If successful, a *memory/* folder will contain both the physical memory
      dump as well as some metadata.

## Encryption

It might be the case that carrying the acquisitions unencrypted might expose
yourself, and even more so those you acquired data from, to significant risk.
For example, you might be stopped at a problematic border and your Snoopdigg
drive could be seized. The raw data might not only expose the purpose of your
trip, but it will likely contain very sensitive data (particularly in the memory
image).

Ideally you would have the drive fully encrypted, but because of practicality
that might not be possible. You could also consider placing Snoopdigg inside a
[VeraCrypt](https://www.veracrypt.fr/) and carry with it a copy of VeraCrypt
to mount it. However, this might be used to force you to unlock and mount it.

Alternatively, Snoopdigg allows to encrypt each acquisition with a provided PGP
public key. Preferably, this public key belongs to a keyset for which you don't
possess or don't carry the private key.

If you place a file called `public.asc` in the same folder as the Snoopdigg
executable, Snoopdigg will automatically attempt to compress and encrypt each
acquisition and delete the original unencrypted copies. Bear in mind, it is
always possible that at least some portion of the unencrypted data could be
recovered through advanced forensics techniques - although we're working to
mitigate that.

## Known issues

The memory acquisition does not work on Windows XP.

## Contacts

Snoopdigg was developed by Claudio "nex" Guarnieri. I can be reached at:

    nex@nex.sx
    PGP ID: 0xD166F1667359D880
    PGP Fingerprint: 0521 6F3B 8684 8A30 3C2F  E37D D166 F166 7359 D880

Or alternatively at:

    nex@amnesty.org
    PGP ID: 0x8F28F25BAAA39B12
    PGP Fingerprint: E063 75E6 B9E2 6745 656C  63DE 8F28 F25B AAA3 9B12
