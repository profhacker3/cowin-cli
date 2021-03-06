
## Command line  tool to List and Book Vaccine
cowin-cli is a simple cli tool to book vaccines and list centers using the COWIN API. It also supports **auto captcha completion** .


>Note: By default cowin-cli will not run continoulsy and monitor slot changes, use bash / batch scripts for that purpose, which can be found [here](#scripts).


## Features
* **Zero dependency** : No neeed to install anything, download precompiled binary and run.
* **Automatic captcha support**: credits to https://github.com/ayushchd
* **Scripting support** : scripts are available for all platforms.
* **Reuse OTP** : session token is written to a text file to reuse it later.
* **Cross platform** : Windows, Linux, macOS, Termux.


- [Installation](#installation)
  - [Install via `go get`](#install-via-go-get)
  - [Download precompiled binaries](#download-precompiled-binaries)
  - [Android Termux](#android-termux)
- [Getting Started](#getting-started)
  - [List vaccine centers](#list-vaccine-centers)
  - [Book Vaccine](#book-vaccine)
  - [Termux Auto OTP](#termux-auto-otp)
  - [Scripts](#scripts)
  - [Options](#options)
    - [List Centers:](#list-center)
    - [Book Vaccine:](#book-vaccine)
- [Known issues](#known-issues)
- [License](#license)


## Installation

### Install via `go get`
```bash
$ go get -u github.com/anoop142/cowin-cli
```
> **Note** : go version 1.16+ is required.

**OR**

### Download precompiled binaries.
Precompiled binaries are avalailable for Windows and linux.
Download them at 
**[Releases](https://github.com/anoop142/cowin-cli/releases)** page.

### Android Termux 
Follow these steps to set up in termux.
```bash
# Install packages
$ pkg i golang git
# Add go bin to PATH
$ echo "export GOPATH=$HOME/go\nexport PATH=$PATH:$GOROOT/bin:$GOPATH/bin" >> ~/.bashrc
$ source ~/.bashrc
#  Install cowin-cli
$ go get -u github.com/anoop142/cowin-cli
```


## Getting Started
There are two modes

* List mode
* Booking mode

### **List vaccine centers**

```
cowin-cli -s state -d district [-v vaccine1,vaccine2] [-m age] [-i] [-b]  [-c dd-mm-yyyy] [-dose dose]
```
### Example 1
```console
$ cowin-cli -s kerala -d alappuzha 

Thazhakara PHC
Kayamkulam THQH  
```

### Example 2
```console
$ cowin-cli -s kerala -d alappuzha -i -m 45 -v "covaxin,covishield" -b -dose 1

Kalavoor PHC  Free  18-05-2021  11 COVAXIN 45 Dose-1
Vandanam MCH  Free  18-05-2021  4 COVISHIELD 45 Dose-1
Mannanchery PHC  Free  18-05-2021  7 COVISHIELD 45 Dose-1
```

The `-i` option displays all extra info like date, vaccine name, age...
`-b'` prints only bookable centers.



### **Book Vaccine**

You can specify mobile number, centers to auto book, age, name etc. 
If not, you will be prompted to enter it appropriately.
```console
$  cowin-cli -sc -state -d district [-no mobileNumber] [-v vaccine1,vaccine2] [-name Name] [-centers center1,cetner2 ] [-slot slotTime] [-ntok]  [-dose dose]
```
### Example 1
```console
$  cowin-cli -sc -s kerala -d alappuzha -no 9123456780

+----+---------------+-----------+---------+-----------+------+
| ID | CENTER        | FREE TYPE | MIN AGE | VACCINE   | DOSE |
+----+---------------+-----------+---------+-----------+------+
|  0 | Aroor FHC     | Free      |      45 | COVISHIELD|   1  |
|  1 | Ala PHC       | Free      |      45 | COVISHIELD|   1  |
|  2 | Kalavoor PHC  | Free      |      45 | COVISHIELD|   2  |
+----+---------------+-----------+---------+-----------+------+

Enter Center ID : 1
Enter OTP : xxxxx

+----+---------------+
| ID |     NAME      |
+----+---------------+
|  0 | John doe      |
|  1 | Jane doe      |
|  2 | All           |
+----+---------------+

Enter name ID : 1

Appointment scheduled successfully!
```
>Note: By default cowin-cli will reuse token, so until the token expires , you don't need to enter otp again.

you can specify most of the details for booking the vaccine

### Example 2
```console
$  cowin-cli -sc -s kerala -d alappuzha -no 9123456780 -name "John doe" -centers "Aroor FHC,Ala PHC" -v "covaxin,sputnik v" -dose 2

Center : Aroor FHC COVAXIN Dose-2
Enter OTP :  xxxxx
```
>**Note**: -centers "any" to auto select any center.
>-name "all" to book for all under same mobile no.
 


### Termux Auto OTP
It's possible to detect OTP message and get OTP in Termux without user input. use -aotp flag to invoke this feature.

You need to first setup termux to read sms.

  1.Install Termux API apk from Fdroid

  2.Install termux-api package 

  ```bash
  # Install termux-api package
  $ pkg i termux-api
  # To give permisiion
  $ termux-list-sms
  # Example
  $ cowin-cli -s kerala -d alappuzha -sc -no 9123456789 -aotp
  ```

### Scripts
Scripts are available for notifying and booking using cowin-cli [here](scripts). You need to edit the vaules of the script like district name, mobile number etc..

### Options

```
  -s	state State Name
  -d  district District name
  -version	Show version
  -h  Show Help
```

#### List Center:

```
  -b	
        show bookable only
  -c string
        date dd-mm-yyyy (default tomorrow's date)
  -i	
        full info
  -p string
        pincode
  -v string
        vaccine names separated by ','
  -m int
        age
  -dose int
            dose type
```

#### Book Vaccine:

```
    -sc
            invoke schedule vaccine mode
    -name string
            registered name           
    -no string
            mobile number
    -centers string separated by ','
            centers to auto select
    -m int
            min age limit
    -slot string
            slot time (FORENOON default)
    -v string
            vaccine names separated by ','
    -ntok
            don't reuse token
    -dose int
            dose type

```

## Known issues
* Random Unauthenticated access error for no reasons.

## License

GPL 3.0

Copyright (c) Anoop S
