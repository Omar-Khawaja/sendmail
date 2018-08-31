This is a small Golang program that uses the
[net/smtp](https://golang.org/pkg/net/smtp/) package to send e-mails. This is a very simple implementation that requires no authentication and can be used to send out simple e-mail alerts.

You will need to make sure Postfix is installed and running on the same node
this program runs on. Your local ISP has probably blocked port 25, so you
will probably want to play around with this program on a node you have set up in
AWS or any other cloud platform.

Depending on your set up, you may need to tweak your Postfix config. If this is
the case, you will probably want to look at **main.cf** and **master.cf** in the
/etc/postfix directory (assuming you are using Linux).
