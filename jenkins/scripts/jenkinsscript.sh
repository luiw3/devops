Name=$1
Age=$3
LastName=$2
Visible=$4

if [ "$Visible" = "true" ]; then
	echo "Hello I am $Name $LastName and I am $Age years old."
else
	echo "Please pass visible=true!!!"
fi
