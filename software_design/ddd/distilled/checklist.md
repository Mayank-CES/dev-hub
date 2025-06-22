# DDD Distilled: Knowledge Quantification Checklist

## Part I: The Fundamentals

### Core Concepts
- [ ] What is the primary goal of Domain-Driven Design (DDD)? What kind of software complexity is it best suited to tackle?
- [ ] How would I explain the difference between the "Domain" and the "Domain Model" to a junior developer?

### The Ubiquitous Language
- [ ] What is a "Ubiquitous Language," and who is responsible for developing it?
- [ ] Why is the Ubiquitous Language considered the foundation of a DDD project?
- [ ] How and where should the Ubiquitous Language be applied in a project?
- [ ] What does it mean to "distill" knowledge from domain experts into a model?

### Problem vs. Solution Space
- [ ] How do I distinguish between the "Problem Domain" and the "Solution"? Can I give an example?
- [ ] What are "Subdomains," and why is it important to identify them early?

## Part II: Strategic Design

### Subdomains
- [ ] What are the three types of Subdomains, and how do they differ in terms of business importance and development strategy?
    - [ ] How do I identify the **Core Domain**? What is its significance?
    - [ ] When is a Subdomain considered **Supporting**?
    - [ ] What qualifies a Subdomain as **Generic**?
- [ ] How does classifying Subdomains help in making strategic decisions (e.g., build vs. buy)?

### Bounded Contexts
- [ ] What is a "Bounded Context"? Why is it considered the most crucial concept in Strategic Design?
- [ ] What is the relationship between a Bounded Context and a Ubiquitous Language? Can one Bounded Context have multiple Ubiquitous Languages?
- [ ] How does a Bounded Context ensure the integrity and consistency of a domain model?
- [ ] What is the difference between a Subdomain (problem space) and a Bounded Context (solution space)?

### Context Mapping
- [ ] What is a "Context Map," and what purpose does it serve in a project?
- [ ] Can I draw a simple Context Map and explain its components?
- [ ] What are the key differences between a **Partnership** and a **Shared Kernel** integration? What are the risks of a Shared Kernel?
- [ ] When would I choose a **Customer-Supplier** relationship, and what are the responsibilities of each party?
- [ ] In what situation would a team be forced into a **Conformist** pattern? What are the downsides?
- [ ] What is an **Anticorruption Layer (ACL)**? Can I describe a scenario where I would implement one?
- [ ] What defines an **Open-Host Service (OHS)**, and how does it differ from a simple API?
- [ ] What is a **Published Language**, and why is it useful for inter-context communication?
- [ ] When is it appropriate for two contexts to go their **Separate Ways**?
- [ ] What are the tell-tale signs of a **Big Ball of Mud**?

## Part III: Tactical Design

### Building Blocks: Entities vs. Value Objects
- [ ] What is the fundamental difference between an **Entity** and a **Value Object**?
- [ ] How is the identity of an Entity determined versus the identity of a Value Object?
- [ ] What are the key characteristics of a Value Object (e.g., immutability, equality)? Can I give three examples?

### Aggregates
- [ ] What is an **Aggregate**? What is its primary purpose?
- [ ] What is an **Aggregate Root**? What are its specific responsibilities?
- [ ] What are the four core rules that govern the design and use of Aggregates?
- [ ] Why is it a best practice to reference other Aggregates only by their identity?
- [ ] What is the reasoning behind the rule "keep Aggregates small"?

### Other Tactical Patterns
- [ ] What is a **Domain Event**, and what role does it play in a domain model? How does it enable eventual consistency?
- [ ] When should I use a **Factory**? What problem does it solve during object creation?
- [ ] What is the purpose of a **Repository**? How does it provide an abstraction over data persistence?
- [ ] Where should the Repository *interface* live versus its *implementation*?
- [ ] When does it make sense to create a **Domain Service**? What are the key characteristics of a well-designed one?