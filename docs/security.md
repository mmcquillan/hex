# Security

Hex allows you to control who can run which rules by using a simple access control list. The list is comma delimited set of either users or channels in slack. For example, if you want to create a private channel called `#ops` and allow anyone invited to that channel to execute commands, you would set the rule with `"acl": "#ops"`. Note that when one runs the `@hex help` command, you only see the list of commands you are allowed to run.

Another feature of the rules is the channel option which guarentees that the output of that rule will end up in the channel specified. This is useful if you have an ACL at a user level which allows the user to direct message requests to Hex and would output to the user's DM channel in addition to a more public channel to keep everyone in the know.

One additional layer of keeping track of who ran what is the Auditing and Auditing File settings in the [configuration](configuration.md) file for Hex. This keeps a record of every user who issues a command through Hex.


