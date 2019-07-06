---
date: "2018-06-12T09:45:46+05:30"
title: "Battery Notifications in i3"
categories:
  - Hacks
tags:
  - Arch Linux
  - i3
---

I am using *i3* window manager for last seven months, and it has been a pleasant and productive experience so far. There were a few hiccups here and there, but that is expected with such minimalistic setups. One thing that I never noticed was the lack of notifications on critical battery levels. For last few months, my laptop battery was discharging to 0% all the time. Probably this proved to be too fatal for my battery. According to this [article](https://lifehacker.com/5875162/how-often-should-i-charge-my-gadgets-battery-to-prolong-its-lifespan), lithium-ion batteries are not expected to go from 100% to 0% frequently. I recently bought a new battery, and I did not want to reduce the lifespan of this battery too. So I decided to setup battery notifications for my i3 setup.

I [found](https://agorf.gr/2016/06/29/low-battery-notification-in-i3wm/) a bash script which shows a notification using *notify-send* when battery charge level reaches or drops below a configured threshold. However, I had to do some additional steps to make this script work on my system.

The first issue was the *lockfile* program which was not installed in my system. I installed it using following command.

{{< highlight bash >}}
sudo apt install procmail
{{< /highlight >}}

The second issue was more difficult to solve. I planned to set up the script to run every minute using *cron*. However, it turns out that cron operates in a very [minimalistic](http://askubuntu.com/a/23438/173003) environment and notify-send requires the presence of some special variables in the environment. These variables are **DBUS_SESSION_BUS_ADDRESS**, **XAUTHORITY** and **DISPLAY**. To provide the values of these variables to the cron environment, I modified the script and sourced a new file `.bat_envs`.

{{< highlight bash >}}
#!/usr/bin/env bash

. /home/yash/.bat_envs

THRESHOLD=15

lock_path='/tmp/battery.lock'

lockfile -r 0 $lock_path 2>/dev/null || exit

acpi_path=$(find /sys/class/power_supply/ -name 'BAT*' | head -1)
charge_now=$(cat "$acpi_path/charge_now")
charge_full=$(cat "$acpi_path/charge_full")
charge_status=$(cat "$acpi_path/status")
charge_percent=$(printf '%.0f' $(echo "$charge_now / $charge_full * 100"
 | bc -l))
message="Battery running critically low at $charge_percent%!"

if [[ $charge_status == 'Discharging' ]] && [[ $charge_percent -le $THRE
SHOLD ]]; then
  /usr/bin/notify-send -u critical "Low battery" "$message"

  current_date_time="`date +%Y%m%d%H%M%S`";
  echo "[BATTERY LOG] = $charge_percent% on $current_date_time"
fi

rm -f $lock_path
{{< /highlight >}}

Read this blog [post](https://agorf.gr/2016/06/29/low-battery-notification-in-i3wm/) to understand how this script works.

As the notify-send requires some special X session environmental variables, we will need a method to provide these variables to notify-send in cron environment. The safest way to get X session related environmental variables is to get them from the environment of a process of the user who is logged on to X. To obtain these variables, the following script will run every time a user logs in and stores these variables in a file `.bat_envs`.

<!--https://unix.stackexchange.com/a/111194 -->

{{< highlight bash >}}
#!/usr/bin/env bash

env_path="$HOME/.bat_envs"

rm -f "${env_path}"
touch "${env_path}"

copy_envs="XAUTHORITY DISPLAY DBUS_SESSION_BUS_ADDRESS"

for env_name in $copy_envs
do
    env | grep "${env_name}" >> "${env_path}";
    echo "export ${env_name}" >> "${env_path}";
done

chmod 600 "${env_path}"
{{< /highlight >}}

To run this script at startup, I added this file to the i3 config file with the following command.

{{< highlight bash >}}
exec --no-startup-id "path to your script"
{{< /highlight >}}

Then at the end of cron file, I added a new entry for the battery monitoring script.

To open cron file:

{{< highlight bash >}}
crontab -e
{{< /highlight >}}

Now add the following line to the end of the file and save the file.

{{< highlight bash >}}
* * * * * bash "path to your script" >> "path to your log file"
{{< /highlight >}}

Replace the *path to your script* (with double quotes) with your script path and the *path to your log file* with a path where you want to save your log file.

Now every minute, this script will be executed and if your battery percent drops below the threshold value, you will be notified with a notification bubble.

I tested this procedure on *Ubuntu 18.04 with i3*. It should work on Arch Linux and other non-Debian distributions also, but the steps might be slightly different due to various reasons. Please comment if you face any issues with the setup.

Thank you for reading the article. Cheers :smile: