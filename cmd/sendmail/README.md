A command-line tool to send email

### Install
```
$ go get github.com/txthinking/xmail/cmd/sendmail
```

### Usage

```
$ sendmail \
    -server smtp.ym.163.com \
    -port 25 \
    -username bot@ym.txthinking.com \
    -password PASSWORD \
    -from bot@ym.txthinking.com \
    -to cloud@txthinking.com:tmp@ym.txthinking.com \
    -subject "Hey boy" \
    -body "Are you fucking with me?" \
    -att /etc/hosts
```
