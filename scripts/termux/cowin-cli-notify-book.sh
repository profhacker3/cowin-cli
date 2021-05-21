#!/bin/bash
# Anoop S
# shell script to notify and book using cowin-cli
# Install notify-send for notifications


COWIN_CLI="./cowin-cli"

# Interval in seconds
T=15
# centers for grep matching
CENTERS_MATCH='VAIJAPUR SDH|LADGAON PHC|BORSAR PHC'
# centers to auto select
CENTERS="VAIJAPUR SDH"
DISTRICT="Aurangabad "
STATE="Maharashtra"
AGE=45
NAME="Mangala Francis Dushing"
NO="8390172690"
# vaccines seperated by ','
VACCINE="covishield"
DOSE=1


schedule(){
	"$COWIN_CLI" -s "$STATE" -d "$DISTRICT" -sc -no "$NO" -name "$NAME" -centers "$CENTERS" -v "$VACCINE" -dose $DOSE -aotp && exit 0 
}


while :
do
	echo "looking for centers.."

	"$COWIN_CLI" -s "$STATE"  -d "$DISTRICT" -m "$AGE" -b -v "$VACCINE" -dose $DOSE

	if (( $? == 0  )) 
	then
		notify
		schedule
	fi

	sleep $T
done
