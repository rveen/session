/*

Package session provides an easy-to-use, extensible and secure HTTP session implementation and management.

Overview

There are 3 key players in the package:

- Session is the (HTTP) session interface. We can use it to store and retrieve constant and variable attributes from it.

- Store is a session store interface which is responsible to store sessions and make them retrievable by their IDs at the server side.

- Manager is a session manager interface which is responsible to acquire a Session from an (incoming) HTTP request, and to add a Session to an HTTP response to let the client know about the session. A Manager has a backing Store which is responsible to manage Session values at server side.

Players of this package are represented by interfaces, and various implementations are provided for all these players.
You are not bound by the provided implementations, feel free to provide your own implementations for any of the players.

Usage

Usage can't be simpler than this. To get the current session associated with the http.Request:

    sess := session.Get(r)
    if sess == nil {
    	// No session (yet)
    } else {
        // We have a session, use it
    }

To create a new session (e.g. on a successful login) and add it to an http.ResponseWriter (let the client know about that):

    sess := session.NewSession()
    session.Add(sess, w)

Let's see a more advanced session creation: let's provide a constant attribute (for the lifetime of the session) and an initial, variable attribute:

    sess := session.NewSessionOptions(&session.SessOptions{
        CAttrs: map[string]interface{}{"UserName": userName},
        Attrs:  map[string]interface{}{"Count": 1},
    })

And to access these attributes and change value of "Count":

    userName := sess.CAttr("UserName")
    count := sess.Attr("Count").(int) // Type assertion, you might wanna check if it succeeds
    sess.SetAttr("Count", count+1)    // Increment count
To remove a session (e.g. logout):

    session.Remove(sess)

Check out the session demo application which shows all these in action:

https://github.com/icza/session/blob/master/session_demo/session_demo.go

*/
package session
