// mailx (https://github.com/txthinking/mailx). Under the MIT license.

// Package mailx is a lightweight SMTP/Mailgun/etc mail sender.
// Implement RFC0821, RFC0822, RFC1869, RFC2045, RFC2821
// Support html body, don't worry that the receiver's
// mail client can't support html, because Mailer will
// send both text/plain and text/html body, so if the
// mail client can't support html, it will display the
// text/plain body.
package mailx
