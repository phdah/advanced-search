# Results

## Overview of the Process

### 1. Data Scraping
- **Process**: Data was scraped from several URLs associated with our public API documentation.
- **Challenges**: This approach allowed for the rapid collection of a large dataset, but the data wasn't fully cleaned, leading to inconsistencies and potential noise in subsequent processes.

### 2. Document Matching with Elasticsearch
- **Process**: The main program takes a user-provided question and uses it as a match query in Elasticsearch to find the top two matching documents.
- **Challenges**: The algorithm struggled with filtering out filler words in the query, which diluted the accuracy of the document matching. This indicates a need for preprocessing the queries to improve match quality.

### 3. LLM Integration for Answer Extraction
- **Process**: The top two documents, along with the original question, are fed into a Large Language Model (LLM) to extract the answer, with the documents serving as mandatory reference sources.
- **Challenges**: While the LLM generally performs as expected, the presence of scattered and overlapping information within the documents sometimes makes it difficult to pinpoint the most relevant answer. For example, multiple definitions of similar concepts (like what constitutes an aircraft) can confuse the model.

## Results

The performance of the system varied. For some questions, the LLM was able to identify the correct data path and produce accurate answers. However, for others, especially those involving more generic API-related queries, the results were less precise. The broad nature of the API and the diversity of topics covered in the documentation contributed to this inconsistency.

## Recommendations for Improvement

1. **Data Structuring**:
   - Implement more rigorous data cleaning and structuring processes. This would involve preprocessing the scraped data to remove irrelevant content and ensure uniformity across documents.

2. **LLM Tuning**:
   - Experiment with LLM parameters and adjust the prompting strategy to enhance the relevance and accuracy of the responses. Fine-tuning the model on a dataset more closely aligned with the specific types of queries anticipated could also help.

3. **Document Filtering**:
   - Introduce a filtering step to limit the documents passed to the LLM to only those that are highly relevant to the query. This could involve refining the Elasticsearch query to better target specific sections of the documentation or implementing additional layers of processing to extract the most pertinent information before passing it to the LLM.

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
bin/advanced-search "What webhooks exists?" Model response: According to the
Avinode TripManager documentation, there are two types of webhooks:

1. **Avinode Webhooks**: These webhooks are used to notify your application
   about certain events in the Avinode system.
2. **Schedaero Webhooks**: These webhooks are used to notify your application
   about certain events in the Schedaero system.

Additionally, it mentions that there are two types of webhooks for each system:

1. **Avinode Webhooks**:

- A notification is sent when a new request (RFQ) is created.

2. **Schedaero Webhooks**:

- Not mentioned explicitly, but assumed to be similar to Avinode's webhooks with
  notifications about certain events in the Schedaero system.

It also mentions that deep links are used for authentication and authorization
of users when accessing Avinode or Schedaero APIs through a custom application.
```

## Multi questions

```bash
bin/advanced-search "What does an aircraft refer to, and what types exists?"
Model response: According to the Avinode glossary:

**Aircraft**: A vehicle that flies through the air. This can be a fixed-wing or
rotary-winged aircraft.

**Aircraft Category**: A categorization of aircraft based on their size and
performance capabilities (e.g., Light, Medium, Heavy).

**Aircraft Type**: A specific type of aircraft, such as an Airbus 320 or a
Boeing 737.
```

## Other topis

```bash
bin/advanced-search "What does an aircraft refer to?" Model response: According
to the text, an "aircraft" refers to a specific type of object or entity that
can be referenced in a payload (data) submitted in a request. The context
suggests that it's likely referring to an airplane or other type of aircraft
that has been registered and is being managed within a system called Avinode.

Model response: According to the text, an ultramarathon can be any distance
longer than a traditional marathon (which is 42.195 kilometers or 26.2 miles).
However, some common distances for ultramarathons include:

- 50 km
- 50 miles
- 100 km

So, the answer is: it depends on the specific event, but typically an ultra
marathon is at least 50-100 kilometers or 31-62 miles long!
```
