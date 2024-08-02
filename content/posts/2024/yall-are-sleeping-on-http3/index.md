---
categories: ["article"]
tags: ["http", "http2", "http3", "web", "webdev"]
date: "2024-08-06T10:00:00Z"
description: ""
cover: "cover.jpg"
images: ["/posts/yall-are-sleeping-on-http3/cover.jpg"]
featured: ""
featuredalt: ""
featuredpath: "date"
linktitle: ""
title: "Y'all are sleeping on HTTP/3"
slug: "yall-are-sleeping-on-http3"
type: "posts"
devtoSkip: true
canonical_url: https://kmcd.dev/posts/yall-are-sleeping-on-http3/
draft: true
---

## Wake up call
This post is about HTTP/3 and QUIC. If you don't know what that is, there are [many](https://www.cloudflare.com/learning/performance/what-is-http3/), [many](https://http.dev/3), [many](https://datatracker.ietf.org/doc/html/rfc9114), [many](https://nordvpn.com/blog/what-is-quic-protocol/), [many](https://blog.cloudflare.com/the-road-to-quic), [many](https://datatracker.ietf.org/doc/rfc9000/), [many](https://www.debugbear.com/blog/http3-quic-protocol-guide) good resources that will get you up to speed. I'm writing this post to enlighten people on what has been happening in the last few years.

**[All major browsers](https://caniuse.com/http3)** support HTTP/3 now.

**[Most major cloud providers](/posts/yall-are-sleeping-on-http3/#cloud-providers)** support HTTP/3 now.

**[Most major load balancers](/posts/yall-are-sleeping-on-http3/#load-balancers)** support HTTP/3 now.

In a short few years **[over 30% of web traffic is served with HTTP/3](https://w3techs.com/technologies/details/ce-http3)**.

{{< bigtext >}}***HTTP/3 isn't the future. It's the present.***{{< /bigtext >}}

Every time I mention HTTP/3 there's always someone who pops up who's completely unaware that it even exists, thinks that it's some minor change or thinks that supporting HTTP/3 is cargo cult behavior. If that last point didn't make sense to you, good. I think the phrase "cargo cult" has been so misused in the tech industry that it has become useless. Anyway, every web developer should probably know that HTTP/3 exists because HTTP/3 is a giant change. HTTP/3 abandons TCP in favor of a channel-aware UDP-based protocol called QUIC. To me, ***this feels important!*** This feels like people need to be talking about it, doing more experiments around QUIC, and writing more tooling, security analysis and benchmarks. QUIC has the potential to dethrone TCP as the reliable layer 4 protocol.

Why should you care about HTTP/3 if you're not a big nerd who likes learning about network protocols? It promises to bring a lot of advantages: faster page loads, smoother video streaming, and more resilient connections. Plus, I think web developers need to know about the foundations of their profession. And this foundation has been changing, rapidly.

{{< image src="thinking.png" width="600px" class="center" >}}

## What's wrong with TCP?
### Multiple streams: One Connection
There are issues with multiplexing multiple streams onto one TLS-backed connection. This is because of TCP's ordering guarantee. It doesn't know about the higher-level protocol's streams so data delivery can be blocked simply because TCP is waiting on packets belonging to an irrelevant stream. This is the so-called [head-of-line blocking](https://http3-explained.haxx.se/en/why-quic/why-tcphol) problem that still exists with HTTP/2 because it is built on top of TCP. Why is this important for the web? Because new HTTP/2 connections are fairly expensive and require multiple round-trips to negotiate. You need a round trip for TCP and two round trips for TLS.

### It isn't good in dynamic network environments
As mentioned in the last section, TCP connections are slow to set up because of the number of round trips required. In environments that change often, causing your web clients to switch networks (and client IP addresses), the clients will have to re-establish an entirely new connection each time you switch networks. Think about the large pause that happens when you switch wifi networks while using video chat. That's exactly this issue. TCP was not designed to handle this situation without that pause. In the following section, I'm going to tell you how QUIC (and HTTP/3) handles this situation in a much more reliable and smooth way.

## Enter: QUIC and HTTP/3
### Faster Connections
First off, QUIC requires far fewer round trips to set up an encrypted connection. Instead of 3 needed for TLS+HTTP/2, QUIC just has one.

### Zero Round Trip Time (0-RTT) Resumption
[QUIC connections can be resumed](https://blog.apnic.net/2023/10/02/how-quic-helps-you-seamlessly-connect-to-different-networks/), even if the client's IP address changes. This allows connections to be re-established with ***zero*** round trips, so when you are on a conference call and switch wifi networks your HTTP/3 connection can resume instantaneously. There have been some security concerns around this feature, but there have been clever solutions to that problem as well.

### Multiplexing
Unlike TCP+HTTP/2, multiple streams of data can be sent over a single connection concurrently without blocking each other because the guaranteed ordering happens at a stream level, not for the entire connection. This completely fixes the head-of-line blocking issue mentioned above.

### Improved Congestion Control
QUIC's [more responsive congestion control](https://www.catchpoint.com/http2-vs-http3/quic-vs-tcp) leads to faster recovery from packet loss.

### "But I heard UDP is unreliable"
It's a common misconception that UDP is inherently unreliable compared to TCP. While it's true that UDP doesn't offer the same built-in guarantees as TCP, QUIC implements those same guarantees on top of UDP.

Think of it like this: TCP is like a delivery service that meticulously tracks every package and resends any that go missing. UDP, on the other hand, is more like a bulk mailer, sending out a flood of letters and hoping most of them arrive. However, QUIC builds its own tracking system on top of UDP, ensuring reliable delivery of data.

This means QUIC can enjoy UDP's speed and flexibility without sacrificing reliability. In fact, in some cases, QUIC can even outperform TCP in terms of packet loss recovery, thanks to its advanced congestion control mechanisms.

So, while the "UDP is unreliable" mantra might have been true in the past, protocols built on UDP can be even more reliable than TCP.

## Let's see how far we've come
Okay, so I mentioned how available HTTP/3 is a few times now. But let's look at the specifics: browser support, cloud support and support when self-hosting.
{{< image src="tech-elite.png" width="600px" class="center" >}}

### Web Browsers
Let's start with web browsers. HTTP/3 support doesn't matter if browsers don't support the technology. [CanIUse](https://caniuse.com/http3) shows the support for each browser and it looks very good for all major browsers:

- [Chrome](https://www.chromium.org/quic/)
- [Firefox](https://hacks.mozilla.org/2021/04/quic-and-http-3-support-now-in-firefox-nightly-and-beta/)
- [Edge](https://techcommunity.microsoft.com/t5/discussions/how-to-enable-http3/m-p/2265230)
- Opera

The only caveat is that Safari only enables HTTP/3 for a portion of users, but Apple will come around eventually.

### Cloud Providers
Next, if you want to host websites with HTTP/3 it would be awesome if cloud providers supported HTTP/3. Well, you're in luck. Most of them do:

- [Cloudflare](https://developers.cloudflare.com/speed/optimization/protocol/http3/)
- [Google Cloud CDN and Load Balancer](https://cloud.google.com/blog/products/networking/cloud-cdn-and-load-balancing-support-http3)
- [AWS CloudFront](https://aws.amazon.com/about-aws/whats-new/2022/08/amazon-cloudfront-supports-http-3-quic/)
- [Akamai CDN](https://techdocs.akamai.com/property-mgr/docs/http3-support)
- [Azure Application Gateway](https://techcommunity.microsoft.com/t5/azure-networking-blog/quic-based-http-3-with-application-gateway-feature-information/ba-p/3913972)
- [CDN77](https://www.cdn77.com/blog/gquic-support-road-to-http3)
- [Fastly CDN](https://docs.fastly.com/en/guides/enabling-http3-for-fastly-services)

This list is not exhaustive, obviously, but

### Load Balancers
But what if you're setting up your own infrastructure, foregoing the cloud? What are the available options?

- [nginx](https://nginx.org/en/docs/quic.html)
- [Envoy](https://gateway.envoyproxy.io/latest/tasks/traffic/http3/)
- [Caddy](https://caddyserver.com/docs/caddyfile/options#section-global-options)
- [LiteSpeed](https://docs.litespeedtech.com/lsws/cp/cpanel/quic-http3/)
- [H2O](https://h2o.examp1e.net/configure/http3_directives.html)

### Is HTTP/2 already dying?
Take a look at the comparative usage of HTTP/2 and HTTP/3 over the last few years:

{{< chart >}}
{
  type: 'line',
  options: {
    plugins: {
        title: {
            display: true,
            text: 'Usage Statistics of HTTP/2 and HTTP/3'
        }
    }
  },
  data: {
    labels: [
        '2016',
        '2017',
        '2018',
        '2019',
        '2020',
        '2021',
        '2022',
        '2023',
        '2024',
        '2024 (Aug)',
    ],
    datasets: [{
      label: 'HTTP/2',
      data: [5.6, 11.2, 23.1, 32.5, 42.6, 49.9, 46.9, 40.0, 35.5, 35.4],
      fill: false,
      borderColor: 'rgb(75, 192, 192)',
      tension: 0.1
    },
    {
      label: 'HTTP/3',
      data: [null, null, null, null, 0, 2.3, 4.1, 24.2, 25.1, 27.8, 30.8],
      fill: false,
      borderColor: 'rgb(255, 99, 132)',
      tension: 0.1
    }]
  }
}
{{< /chart >}}
[Source: w3techs.com](https://w3techs.com/technologies/history_overview/site_element/all/y)

From the graph below you can see that in a short few years, HTTP/3 usage is rapidly approaching the same usage as HTTP/2. We had "peak HTTP/2" in 2021 and maybe next year we will see HTTP/3 overtake HTTP/2 to be the de-facto standard for new web deployments. If I were to guess, I think HTTP/2 might be deprecated sooner than HTTP/1.1.

{{< image src="reaper.png" width="600px" class="center" >}}

I will be honest though, different data aren't as amazing for HTTP/3 from [Cloudflare](https://blog.cloudflare.com/http3-usage-one-year-on) and [Internet Society Pulse](https://pulse.internetsociety.org/technologies). Both of those data sources show HTTP/2 still handling the majority of requests.

## Challenges Ahead
There are two main areas to focus on with QUIC: adding more tooling and language support for the protocol.

### Tooling/Language Support
Even though browsers and load balancers have good support for QUIC, most programming languages don't support HTTP/3 because QUIC presents a vastly different way to communicate. Without kernel support, adding QUIC to a language is a bit like re-implementing TCP in the language. TCP is usually relatively easy to add because OS kernels typically implement TCP for you and provide bindings. As far as I know, that isn't the case for QUIC.

### Will QUIC stay in userspace?
From what I've seen, there's no well-supported kernel module for QUIC. There exist some projects that do [add a kernel module for Linux](https://github.com/lxin/quic) but it doesn't seem like it's heavily used yet. TCP has benefitted a lot from being backed by OS kernels, which gives it performance optimizations not available in user space but that strength is also a weakness. TCP is now hard to change and adapt. It appears like QUIC has evolved much quicker because of its user-space implementations, but should it stay there? Should there be [a hybrid approach](https://lwn.net/Articles/965134/)? To me, those are open questions and I'm curious how this will unfold over time.

## The Future is QUIC
HTTP/3 and QUIC are more than just incremental improvements; they represent a fundamental shift in how we build the web. By overcoming the limitations of TCP, QUIC offers a faster, more reliable, and more secure foundation for modern Internet communication.

The benefits are clear: faster page loads, smoother video streaming, more responsive applications, and seamless transitions between networks. This translates to a better user experience, increased engagement, and ultimately, a more vibrant and dynamic internet.

But QUIC's potential extends far beyond the web. Its ability to handle real-time communication, coupled with its built-in security features, makes it an attractive option for applications like online gaming, video conferencing, and more.

While the transition to HTTP/3 is well underway, it's far from complete. There are still challenges to overcome, such as improving tooling, expanding language support, and ensuring compatibility with existing infrastructure.

However, the momentum is undeniable. With major browsers, cloud providers, and load balancers embracing QUIC, the future is bright. As developers and engineers continue to innovate and build on this foundation, we can expect to see even more exciting possibilities emerge.

So, what are you waiting for? If you're a web developer or tech enthusiast, now is the time to dive into QUIC and start exploring its potential. Experiment with new applications, build innovative tools and share your findings with the community. [I've started doing this by trying HTTP/3 with gRPC](/posts/grpc-over-http3/). The future of the internet is being written in QUIC, and you can be a part of it.

Here's a link dump of more resources on this:
- https://www.fastly.com/blog/measuring-quic-vs-tcp-computational-efficiency/
- https://github.com/lxin/quic
- https://www.chromium.org/quic/
- https://codilime.com/blog/http3-protocol/
- https://http3-explained.haxx.se/en
- https://caniuse.com/http3
- https://sanjeev41924.medium.com/http-3-challenges-to-security-and-possible-response-e81f429506e0
- https://pulse.internetsociety.org/blog/the-challenges-ahead-for-http-3
- https://medium.com/tech-internals-conf/http-3-shiny-new-thing-or-more-issues-6e4fe14e52ea
- https://w3techs.com/technologies/history_overview/site_element/all/y
- https://blog.apnic.net/2023/09/25/why-http-3-is-eating-the-world/
- https://blog.apnic.net/2023/10/02/how-quic-helps-you-seamlessly-connect-to-different-networks/
- https://lwn.net/Articles/965134/