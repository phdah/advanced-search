# Results

## Single question

```bash
bin/advanced-search "What does an aircraft refer to?"
Model response:
 According to the text, an "aircraft" refers to a specific type of object or
entity that can be referenced in a payload (data) submitted in a request. The
context suggests that it's likely referring to an airplane or other type of
aircraft that has been registered and is being managed within a system called
Avinode.
```

```bash
bin/advanced-search "What is a deep link?"
Model response:
 Working-with-Deep-Links: "Some resources in Avinode and Schedaero APIs may
include URIs to a corresponding page in the Avinode/Schedaero application,
often known as a “deep link”."
```

```bash
bin/advanced-search "What webhooks exists?"
Model response:
 According to the Avinode TripManager documentation, there are two types of webhooks:

1. **Avinode Webhooks**: These webhooks are used to notify your application about certain events in the Avinode system.
2. **Schedaero Webhooks**: These webhooks are used to notify your application about certain events in the Schedaero system.

Additionally, it mentions that there are two types of webhooks for each system:

1. **Avinode Webhooks**:
 * A notification is sent when a new request (RFQ) is created.
2. **Schedaero Webhooks**:
 * Not mentioned explicitly, but assumed to be similar to Avinode's webhooks with notifications about certain events in the Schedaero system.

It also mentions that deep links are used for authentication and authorization of users when accessing Avinode or Schedaero APIs through a custom application.
```

## Multi questions

```bash
bin/advanced-search "What does an aircraft refer to, and what types exists?"
Model response:
 According to the Avinode glossary:

**Aircraft**: A vehicle that flies through the air. This can be a fixed-wing or rotary-winged aircraft.

**Aircraft Category**: A categorization of aircraft based on their size and performance capabilities (e.g., Light, Medium, Heavy).

**Aircraft Type**: A specific type of aircraft, such as an Airbus 320 or a Boeing 737.
```
