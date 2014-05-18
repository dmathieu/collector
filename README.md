# Collector
[![Build Status](https://travis-ci.org/dmathieu/collector.svg?branch=master)](https://travis-ci.org/dmathieu/collector)

## Why?

On a side project, I wanted to track data from javascript. Ajax calls, DOM rendering time, ...
Much like [analytics at github](http://johnnunemaker.com/analytics-at-github/), but in a much simpler way,
because this is a hobby project and I don't have as much traffic as github.

### Installing

Collector is meant to run on heroku. Installing it is pretty easy.

```
heroku create
heroku config:set BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git#go1.2
heroku config:set LIBRATO_EMAIL=<your librato email>
heroku config:set LIBRATO_TOKEN=<your librato token>
git push heroku
```

That's all!

### Usage

I'm using Collector with Javascript only (when using it on a server-side app, I can just make calls to librato. No need for this).  
Here's the small method I'm using:

```
window.collector = {
  host: host,
  collect: function(kind, key, value) {
    var host;
    host = 'http://<app_name>.herokuapp,com/';
    return new Ember.RSVP.Promise(function(resolve, reject) {
      return jQuery.ajax({
        url: "" + host + "/collect",
        method: "POST",
        data: {
          metric: {
            key: key,
            value: value,
            kind: kind
          }
        },
        success: function(r) {
          return resolve(r);
        },
        error: function(r) {
          return reject(r);
        }
      });
    });
  }
};
```

This method can be called like this:

    collector.collect('meter', 'app.pageview', 1)

Which will send an `app.pageview` metric with the value '1' to your collector instance.
This data will be sent to librato, where you can view it afterwards.

That's all, nothing complicated. Though you should be able to build powerful javascript frameworks to send more data if you need to.

## Contributing

If you think Collector is great but can be improved, feel free to contribute. To do so, you can :

* [Fork](http://help.github.com/forking/) the project
* Do your changes and commit them to your repository
* Test your changes. We won't accept any untested contributions (except if they're not testable).
* Create an [issue](https://github.com/dmathieu/collector/issues) with a link to your commits.

And that's it! I'll soon take a look at your issue and review your changes.

## Author and Credits

Damien MATHIEU :: 42 (AT|CHEZ) dmathieu.com
