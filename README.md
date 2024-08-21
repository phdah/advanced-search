# advanced-search
Hackathon project

## Description
A problem that has been expressed from the operations side of the organisation is that there is a lot of documentations that is hard to navigate and find in Confluence. Other companies has built internal bot that can answers company specific internal questions. These model I think is from OpenAI and has been trained on their own data to be able to answer questions.

With the release of open source models such as `llama3` and many more, the option of running self hosted LLMs to build custom RAGs is within reach, and without the risk of exposing/sharing data outside of the organisation. Hence, I have been thinking about how effective it would be to use an out of the box LLM to answer questions based on documents provided as part of the prompted question.

This could probably be achieved by combining
1. Good data
  - accessible through some structured approach (Elastic Search)
2. An open source LLM agent

![image](./img/advanced_search.jpg)

## Steps
- [x] Plan
  - Review problem statement and existing challenges with document navigation in Confluence.
  - Initial design ideas sketched.

- [ ] Pick documents
  - Identify key documents and data sources that need to be indexed. -> [Avinode api](https://developer.avinodegroup.com/docs/introduction) which will have to be scraped.
  - Gather necessary documentation from some source (wikipedia or pip?) and other relevant sources.

- [ ] Setup elastic search
  - Install and configure Elasticsearch for document indexing and search capabilities. [insta-infra](https://github.com/data-catering/insta-infra)
  - Design and implement the structure for document indexing.

- [ ] Setup LLM query builder
  - Integrate an open-source LLM (e.g., llama3(.1)) with the system.
  - Develop a query builder to interact with the LLM, leveraging the indexed documents.

- [ ] Setup query execution and response handler
  - Implement logic to execute queries against the LLM using the structured data from Elasticsearch.
  - Create a response handler to parse and present the answers effectively.

- [ ] Testing and optimization
  - Conduct tests to evaluate the system's accuracy and responsiveness.
  - Fine-tune the LLM and Elasticsearch configurations for optimal performance.

- [ ] Future improvements
  - Explore the possibility of training a custom model on Avinode data for more accurate responses.
  - Investigate additional tooling and structures for enhancing search capabilities.
