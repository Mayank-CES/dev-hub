Here's a more concise summary of Chapter 4, "Strategic Design with Context Mapping," in Markdown format, with definitions and examples clearly segregated.

# Chapter 4: Strategic Design with Context Mapping

[cite\_start]Previous chapters established the role of *Bounded Contexts* in DDD[cite: 1637]. [cite\_start]This chapter focuses on *Context Mapping*, which describes the integration between these Bounded Contexts[cite: 1238].

## Definitions

  * [cite\_start]**Context Mapping**: Defines team relationships and technical mechanisms for integration between two Bounded Contexts[cite: 1264]. [cite\_start]It represents the translation between their respective *Ubiquitous Languages*[cite: 1254].

    ```mermaid
    graph LR
        BC1(Bounded Context 1) -- "Mapping Line" --> BC2(Bounded Context 2)
    ```

    *Figure: Basic Context Mapping*

  * [cite\_start]**Partnership**: Two teams, each owning a Bounded Context, align on dependent goals, succeeding or failing together[cite: 1273, 1275]. [cite\_start]Requires frequent synchronization and continuous integration due to high commitment[cite: 1276, 1278].

    ```mermaid
    graph LR
        subgraph Team 1
            style Team 1 fill:#add8e6,stroke:#333,stroke-width:2px
            BC1(Bounded Context 1)
        end
        thick_line_style { stroke-width: 4px; }
        subgraph Team 2
            style Team 2 fill:#add8e6,stroke:#333,stroke-width:2px
            BC2(Bounded Context 2)
        end
        BC1 --o BC2;
        linkStyle 0 stroke-width:4px,fill:none,stroke:#000;
    ```

    *Figure: Partnership Context Mapping*

  * [cite\_start]**Shared Kernel**: Two or more teams share a small, common model[cite: 1284]. [cite\_start]Requires open communication and constant agreement on shared elements[cite: 1287].

    ```mermaid
    graph LR
        subgraph Team 1
            style Team 1 fill:#add8e6,stroke:#333,stroke-width:2px
            BC1_P(BC 1 Parts)
            SK1((Shared Kernel))
        end
        subgraph Team 2
            style Team 2 fill:#add8e6,stroke:#333,stroke-width:2px
            BC2_P(BC 2 Parts)
            SK2((Shared Kernel))
        end
        SK1 ~~~ SK2
        linkStyle 0 stroke-dasharray: 5 5
        style SK1 fill:#99ccff,stroke:#000,stroke-dasharray: 5 5
        style SK2 fill:#99ccff,stroke:#000,stroke-dasharray: 5 5
    ```

    *Figure: Shared Kernel Context Mapping*

  * [cite\_start]**Customer-Supplier**: A *Supplier* (upstream) provides what a *Customer* (downstream) needs[cite: 1294, 1295]. [cite\_start]Supplier dictates the terms, but should remain responsive to Customer needs[cite: 1296, 1297].

    ```mermaid
    graph LR
        subgraph Team 1
            style Team 1 fill:#add8e6,stroke:#333,stroke-width:2px
            BC1_U(Upstream)
            UM[U]
        end
        subgraph Team 2
            style Team 2 fill:#add8e6,stroke:#333,stroke-width:2px
            BC2_D(Downstream)
            DM[D]
        end
        UM --x BC2_D
        BC1_U --UM---DM
        linkStyle 0 stroke-width:2px,fill:none,stroke:#000;
        linkStyle 1 stroke-width:2px,fill:none,stroke:#000;
        style UM fill:transparent,stroke:transparent;
        style DM fill:transparent,stroke:transparent;
    ```

    *Figure: Customer-Supplier Context Mapping*

  * [cite\_start]**Conformist**: Downstream team conforms to the upstream model's *Ubiquitous Language* because the upstream has no motivation to support their specific needs, and translation is not feasible[cite: 1303, 1304].

    ```mermaid
    graph LR
        subgraph Upstream Team
            style Upstream Team fill:#add8e6,stroke:#333,stroke-width:2px
            BC_Up(Bounded Context - Upstream)
            UM[U]
        end
        subgraph Downstream Team
            style Downstream Team fill:#add8e6,stroke:#333,stroke-width:2px
            BC_Down(Bounded Context - Downstream)
            DM[D]
        end
        UM --x BC_Down
        BC_Up --UM---DM
        linkStyle 0 stroke-width:2px,fill:none,stroke:#000;
        linkStyle 1 stroke-width:2px,fill:none,stroke:#000;
        style UM fill:transparent,stroke:transparent;
        style DM fill:transparent,stroke:transparent;
    ```

    *Figure: Conformist Context Mapping*

  * [cite\_start]**Anticorruption Layer (ACL)**: A defensive translation layer created by the downstream team to isolate its model from an upstream model's *Ubiquitous Language*[cite: 1313, 1314]. [cite\_start]It protects the downstream model from "foreign concepts"[cite: 1316].

    ```mermaid
    graph LR
        subgraph Team 1 (Upstream)
            style Team 1 (Upstream) fill:#add8e6,stroke:#333,stroke-width:2px
            U1[U]
        end
        subgraph Team 2 (Downstream)
            style Team 2 (Downstream) fill:#add8e6,stroke:#333,stroke-width:2px
            ACL[ACL]
            D1[D]
        end
        U1 --- ACL
        ACL --- D1
    ```

    *Figure: Anticorruption Layer*

  * [cite\_start]**Open Host Service (OHS)**: Defines a well-documented, "open" protocol or interface to access a Bounded Context's services[cite: 1322, 1323, 1324]. [cite\_start]Facilitates easy integration for other contexts[cite: 1323].

    ```mermaid
    graph LR
        subgraph Team 1
            style Team 1 fill:#add8e6,stroke:#333,stroke-width:2px
            BC1(Bounded Context 1)
            OHS[OHS]
        end
        subgraph Team 2
            style Team 2 fill:#add8e6,stroke:#333,stroke-width:2px
            BC2(Bounded Context 2)
        end
        BC1 -- OHS --> BC2
    ```

    *Figure: Open Host Service*

  * [cite\_start]**Published Language (PL)**: A well-documented information exchange language for easy consumption and translation by consuming Bounded Contexts[cite: 1331, 1332]. [cite\_start]Often served by an Open Host Service[cite: 1334].

    ```mermaid
    graph LR
        subgraph Team 1
            style Team 1 fill:#add8e6,stroke:#333,stroke-width:2px
            BC1(Bounded Context 1)
            PL[PL]
        end
        subgraph Team 2
            style Team 2 fill:#add8e6,stroke:#333,stroke-width:2px
            BC2(Bounded Context 2)
        end
        BC1 -- PL --> BC2
    ```

    *Figure: Published Language*

  * [cite\_start]**Separate Ways**: Integration is not pursued if the payoff from consuming another *Ubiquitous Language* is insignificant or the functionality is incomplete[cite: 1339, 1340]. [cite\_start]A specialized solution is built instead[cite: 1341].

    ```mermaid
    graph LR
        subgraph Team 1
            style Team 1 fill:#add8e6,stroke:#333,stroke-width:2px
            BC1(Bounded Context 1)
        end
        subgraph Team 2
            style Team 2 fill:#add8e6,stroke:#333,stroke-width:2px
            BC2(Bounded Context 2)
        end
        BC1 --X BC2
    ```

    *Figure: Separate Ways Context Mapping*

  * [cite\_start]**Big Ball of Mud (in Context Mapping)**: Existing complex, tangled legacy systems[cite: 1343, 1351]. [cite\_start]Integration leads to cross-contamination, "whack-a-mole" issues, and reliance on tribal knowledge[cite: 1345, 1346, 1347]. [cite\_start]If forced to integrate, an *Anticorruption Layer* is crucial to protect your model[cite: 1353].

## Examples and Integration Mechanisms

Integration interfaces depend on the providing Bounded Context. [cite\_start]Common types include RPC via SOAP, RESTful HTTP, or messaging[cite: 1378, 1379]. [cite\_start]Database or file system integration should generally be avoided; if unavoidable, an *Anticorruption Layer* is strongly recommended[cite: 1380].

### RPC with SOAP

  * [cite\_start]**Description**: Makes remote service calls appear local[cite: 1397]. [cite\_start]Uses network transport for requests and responses[cite: 1398].
    ```mermaid
    graph LR
        ServiceBC(Service Bounded Context) -- "SOAP (XML)" --> ClientBC(Client Bounded Context)
    ```
    *Figure: RPC with SOAP*
  * [cite\_start]**Challenges**: Prone to network failures, unanticipated latency, and strong coupling[cite: 1399, 1400, 1404].
  * [cite\_start]**Mitigation**: Ideally, the service Bounded Context provides an *Open Host Service* with a *Published Language*[cite: 1414]. [cite\_start]The client Bounded Context should implement an *Anticorruption Layer* to isolate its model[cite: 1415].
    ```mermaid
    graph LR
        ServiceBC(Service Bounded Context) -- OHS/PL --> ClientBC(Client Bounded Context)
        ClientBC -- ACL --> ServiceBC
    ```
    *Figure: RPC with OHS/PL and ACL*

### RESTful HTTP

  * [cite\_start]**Description**: Focuses on resource exchange using standard HTTP operations (POST, GET, PUT, DELETE)[cite: 1425]. [cite\_start]Effective for distributed computing APIs[cite: 1426].
    ```mermaid
    graph LR
        ServiceBC(Service Bounded Context) -- RESTful HTTP --> ClientBC(Client Bounded Context)
    ```
    *Figure: RESTful HTTP Integration*
  * [cite\_start]**Challenges**: Can fail due to network or service provider issues, similar to RPC[cite: 1443].
  * [cite\_start]**Best Practice**: A service Bounded Context should provide an *Open Host Service* and a *Published Language* via its REST interface[cite: 1437, 1438].
  * [cite\_start]**Common Mistake**: Designing REST resources that directly mirror *Aggregates* in the domain model[cite: 1451]. [cite\_start]This forces clients into a *Conformist* relationship[cite: 1452]. [cite\_start]Instead, resources should be "synthetic," designed to match client-driven use cases, reflecting client needs rather than internal model composition[cite: 1453, 1454, 1457].

### Messaging

  * [cite\_start]**Description**: Uses asynchronous communication, often via *Domain Events*, where a publisher notifies interested subscribers[cite: 1464]. [cite\_start]More robust than synchronous methods as it removes temporal coupling[cite: 1465]. [cite\_start]Latency is expected and planned for[cite: 1466, 1527].
    ```mermaid
    graph LR
        PubBC(Publishing Bounded Context) -- "Domain Event" --> MsgMech[Messaging Mechanism]
        MsgMech -- "Domain Event" --> SubBC(Subscribing Bounded Context)
    ```
    *Figure: Messaging Mechanism*
  * **Key Principles**:
      * [cite\_start]*At-Least-Once Delivery*: The messaging mechanism ensures messages are eventually delivered, even if multiple times[cite: 1509, 1515, 1517].
      * [cite\_start]*Idempotent Receiver*: Consumers must be designed to handle duplicate messages correctly, producing the same result regardless of how many times an operation is performed[cite: 1510, 1523, 1524, 1526].
  * [cite\_start]**Domain Event Consumers**: Should rely on the *schema* (their *Published Language*) of events, not the publisher's event types, to avoid *Conformist* relationships[cite: 1488, 1489].
  * [cite\_start]**Commands**: A client Bounded Context can send a *Command Message* to a service Bounded Context to request an action, with the outcome received as a *Domain Event*[cite: 1499, 1500].
  * [cite\_start]**Integration Train Wrecks**: Avoid synchronous, blocking requests between Bounded Contexts as a direct result of handling another blocking request[cite: 1474, 1475]. [cite\_start]Asynchronous messaging prevents this[cite: 1477].

## Example in Context Mapping (Insurance Policy)

  * [cite\_start]**Scenario**: The "policy of record" originates in the *Underwriting Context*[cite: 1539, 1540].
  * [cite\_start]**Integration**: When a `Policy` is issued in the *Underwriting Context*, it publishes a `PolicyIssued` *Domain Event*[cite: 1552]. [cite\_start]Other Bounded Contexts (e.g., Claims, Inspections) subscribe to this event and create corresponding `Policy` components[cite: 1553].
    ```mermaid
    graph LR
        UnderwritingBC(Underwriting Context) -- "publishes PolicyIssued" --> PolicyIssuedEvent[PolicyIssued]
        PolicyIssuedEvent -- "causes creation" --> ClaimsBC(Claims Context)
        PolicyIssuedEvent -- "causes creation" --> InspectionsBC(Inspections Context)
    ```
    *Figure: PolicyIssued Event Flow*
  * [cite\_start]**Traceability and Querying**: The `PolicyIssued` event contains the official `policyId`[cite: 1562]. [cite\_start]Subscribing contexts store this `policyId` (e.g., as `issuedPolicyId`) for traceability[cite: 1563, 1564]. [cite\_start]If more data is needed, the subscribing context can query back on the *Underwriting Context* using the `issuedPolicyId`[cite: 1564, 1565]. [cite\_start]This query could use a RESTful *Open Host Service* and *Published Language* from the *Underwriting Context*[cite: 1581].

### Enrichment versus Query-Back Trade-offs

  * [cite\_start]**Enrichment**: Including sufficient data in *Domain Events* for all consumers' needs allows greater consumer autonomy[cite: 1567, 1569]. [cite\_start]However, predicting all future data needs is difficult, and excessive enrichment can obscure meaning or pose security risks[cite: 1570, 1571].
  * **Query-Back**: Keeping *Domain Events* thin and allowing consumers to request additional data when needed. [cite\_start]This is preferable for security-sensitive data or when consumer needs are varied and unpredictable[cite: 1568, 1572].
  * [cite\_start]A balanced blend of both approaches is often optimal[cite: 1573].

## Summary

This chapter provided:

  * [cite\_start]An understanding of various *Context Mapping* relationships (Partnership, Customer-Supplier, Conformist, Anticorruption Layer, Open Host Service, Published Language, Separate Ways)[cite: 1622].
  * [cite\_start]Guidance on using *Context Mapping* for integration with RPC, RESTful HTTP, and messaging[cite: 1623].
  * [cite\_start]Insights into how *Domain Events* function with messaging[cite: 1624].
  * [cite\_start]A foundation for building Context Mapping expertise[cite: 1625].

[cite\_start]For thorough coverage, refer to Chapter 3 of *Implementing Domain-Driven Design* [IDDD][cite: 1626].

```
```
