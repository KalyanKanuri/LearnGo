# Go LLD - Syllabus, Study Plan, and Progress Tracker

*A practical, coding-first roadmap for learning low-level design from first
principles through production-minded design and end-to-end projects.*

## 1. Learning goals

By the end of this roadmap, you should be able to:

- Model a problem domain by identifying objects, state, behavior, invariants,
  and relationships.
- Assign responsibilities deliberately and choose composition, interfaces,
  and dependencies only when they solve a real problem.
- Apply SOLID and selected design patterns to concrete requirements, rather
  than memorizing definitions.
- Write idiomatic Go with small interfaces, clear package boundaries, explicit
  errors, safe concurrency, and testable components.
- Explain design choices and trade-offs in code reviews and LLD interviews.
- Progress from object-level design to a maintainable application-level
  implementation.

## 2. Learning method

Each concept follows this cycle:

1. **Understand the problem:** Define requirements, constraints, edge cases,
   and invariants.
2. **Model before coding:** Identify domain concepts, responsibilities,
   relationships, and public APIs.
3. **Predict and choose:** Make a design decision and explain the trade-off
   before seeing the solution.
4. **Implement in Go:** Write the code myself from acceptance criteria;
   starter code may contain deliberate design flaws.
5. **Test and review:** Use unit tests, table-driven tests, fakes, and
   concurrency tests where appropriate.
6. **Refactor:** Fix review feedback and explain why the new design is better.
7. **Track mastery:** Mark a topic complete only after the implementation,
   tests, and explanation meet the criteria.

## 3. Curriculum syllabus

### Phase 1 - LLD foundations

| Topic                                  | What to learn                                                        | Coding task / evidence                                                                      |
| -------------------------------------- | -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| Problem decomposition and requirements | Separate functional requirements, constraints, and edge cases.       | Write a requirements checklist for a parking lot.                                           |
| Objects, state, behavior, and identity | Identify domain entities and responsibilities.                       | Model `Vehicle`, `ParkingSpot`, `ParkingFloor`, `ParkingLot`, and `ParkingTicket`.          |
| Abstraction and encapsulation          | Protect state and expose meaningful operations.                      | Prevent callers from directly changing parking-spot occupancy.                              |
| Relationships                          | Association, aggregation, composition, inheritance, and dependency.  | Draw a class diagram and explain ownership and lifetimes.                                   |
| Polymorphism and Go interfaces         | Use behavior-based, small interfaces.                                | Model cars, bikes, and trucks through a `Vehicle` interface where justified.                |
| Composition vs. inheritance            | Prefer composition when behaviors or features combine independently. | Model EV, premium, and covered-spot characteristics without a subclass explosion.           |
| UML and interaction flows              | Class, sequence, and state diagrams.                                 | Diagram the Park and Exit workflows.                                                        |

### Phase 2 - Design principles

| Topic                                    | What to learn                                                      | Coding task / evidence                                                                      |
| ---------------------------------------- | ------------------------------------------------------------------ | ------------------------------------------------------------------------------------------- |
| SRP                                      | One responsibility or reason to change; cohesion.                  | Refactor a `ParkingLot` that also calculates fees, emails receipts, and generates reports.  |
| OCP                                      | Extend behavior without repeatedly editing stable code.            | Add a parking-allocation policy without rewriting the service.                              |
| LSP                                      | Subtypes must preserve the expectations of their abstractions.     | Repair an abstraction whose implementation unexpectedly rejects valid operations.           |
| ISP                                      | Clients should depend on small, relevant interfaces.               | Split an oversized notification or payment interface.                                       |
| DIP                                      | High-level policy should not depend directly on low-level details. | Inject payment and storage dependencies into a service.                                     |
| Coupling, cohesion, DRY, KISS, and YAGNI | Recognize useful abstractions versus overengineering.              | Compare two designs and justify the simpler maintainable one.                               |

### Phase 3 - Dependency and application structure

| Topic                  | What to learn                                                       | Coding task / evidence                                                                 |
| ---------------------- | ------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| Dependency injection   | Constructor injection, composition root, and test seams.            | Swap a real payment provider for a fake in tests.                                      |
| Package design in Go   | Package boundaries, exported APIs, and avoiding import cycles.      | Organize domain, service, and adapter packages.                                        |
| Service layer          | Coordinate application use cases.                                   | Implement Park and Exit without placing workflow in `Vehicle` or `ParkingSpot`.        |
| Repository             | Separate persistence operations from business rules when useful.    | Implement an in-memory repository and a fake for tests.                                |
| Domain models vs. DTOs | Keep API and database representations from leaking into the domain. | Map HTTP request DTOs to domain inputs.                                                |
| Error handling         | Sentinel and wrapped errors, validation, and error ownership.       | Return actionable errors for no spot, incompatible vehicle, and duplicate reservation. |

### Phase 4 - Creational patterns

| Topic            | What to learn                                                            | Coding task / evidence                                                    |
| ---------------- | ------------------------------------------------------------------------ | ------------------------------------------------------------------------- |
| Factory Method   | Choose implementation construction when creation varies.                 | Select a payment provider from configuration.                             |
| Abstract Factory | Create compatible families of related objects; use sparingly.            | Model provider-specific payment components only if justified.             |
| Builder          | Construct complex objects with many optional fields or validation rules. | Build a report/export request with validated options.                     |
| Singleton        | Understand global state, lifecycle, and testability risks.               | Review a singleton cache and refactor if shared global state hurts tests. |
| Prototype        | Copy an existing configured object when copying is meaningful.           | Discuss whether configuration cloning needs a pattern.                    |

### Phase 5 - Structural patterns

| Topic                | What to learn                                                           | Coding task / evidence                                                  |
| -------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| Adapter              | Translate an external API into an application-owned interface.          | Wrap a third-party payment SDK.                                         |
| Decorator            | Add behavior around an implementation without modifying it.             | Add logging or metrics around a service.                                |
| Facade               | Provide a simpler interface to a complex subsystem.                     | Expose a checkout facade over inventory, payment, and order components. |
| Proxy                | Control access or add caching, lazy, or remote behavior.                | Implement a caching or authorization wrapper.                           |
| Composite            | Treat individual objects and groups uniformly.                          | Represent nested product categories or file trees.                      |
| Bridge and Flyweight | Separate independent dimensions; share intrinsic state when worthwhile. | Evaluate each against a concrete use case before implementing.          |

### Phase 6 - Behavioral patterns

| Topic                                                 | What to learn                                                                           | Coding task / evidence                                                              |
| ----------------------------------------------------- | --------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| Strategy                                              | Swap algorithms or behavior at runtime or construction time.                            | Implement nearest-spot and premium-allocation policies.                             |
| Observer / pub-sub                                    | Notify independent subscribers about an event.                                          | Send email, update analytics, and emit an event after order placement.              |
| State                                                 | Make valid behavior depend on explicit lifecycle state.                                 | Implement order transitions with invalid-transition tests.                          |
| Command                                               | Represent an operation as a value.                                                      | Model retryable admin operations or queued actions.                                 |
| Chain of Responsibility                               | Pass a request through independent handlers.                                            | Build an HTTP validation or middleware pipeline.                                    |
| Template Method                                       | Share a fixed algorithm skeleton with variable steps; consider composition first in Go. | Compare a template approach with injected functions or interfaces.                  |
| Iterator, Mediator, Memento, Visitor, and Interpreter | Know the problem each solves and its trade-offs.                                        | Complete short scenario exercises; implement only when a real use case warrants it. |

### Phase 7 - Production concerns

| Topic                                   | What to learn                                                               | Coding task / evidence                                                |
| --------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------- |
| Domain invariants and state transitions | Keep invalid states and transitions from entering the system.               | Prevent double booking and invalid ticket closure.                    |
| Concurrency and race conditions         | Mutexes, atomic operations, ownership, contention, and cancellation.        | Make reservation safe under concurrent requests; run race tests.      |
| Consistency and transaction boundaries  | Handle partial failure and multi-step operations.                           | Recover if reservation succeeds but ticket persistence fails.         |
| Retries, idempotency, and timeouts      | Avoid duplicate effects and unbounded work.                                 | Design an idempotent payment request and bounded retry policy.        |
| Observability                           | Structured logs, metrics, traces, and useful error context.                 | Instrument Park and Exit without coupling business rules to a vendor. |
| Testing strategy                        | Unit, integration, contract, and concurrency tests.                         | Use table-driven tests and fakes; identify what should not be mocked. |
| Refactoring and code review             | Review correctness, simplicity, maintainability, security, and performance. | Fix a deliberately flawed service implementation.                     |

### Phase 8 - Backend application architecture

| Topic                              | What to learn                                                      | Coding task / evidence                                                            |
| ---------------------------------- | ------------------------------------------------------------------ | --------------------------------------------------------------------------------- |
| HTTP handlers and application APIs | Keep transport concerns separate from use cases.                   | Implement REST endpoints for Park, Exit, and ticket lookup.                       |
| Validation and authorization       | Validate at boundaries and enforce permissions at the right layer. | Reject malformed requests and unauthorized admin operations.                      |
| Persistence adapters               | Keep database details outside core business rules.                 | Add a SQL-backed repository behind a narrow interface.                            |
| External integrations              | Timeouts, errors, adapters, and provider-specific details.         | Integrate a mock payment gateway.                                                 |
| Configuration and lifecycle        | Build dependencies at startup and close resources gracefully.      | Wire dependencies in `main` and handle shutdown.                                  |
| Modular monolith                   | Organize modules around cohesive business capabilities.            | Structure an order/inventory/payment application without premature microservices. |

### Phase 9 - End-to-end LLD projects

| Project                     | Concepts                                                                  | Deliverable                                                            |
| --------------------------- | ------------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| Parking Lot                 | Objects, policies, reservation safety, tickets, pricing, entry, and exit. | Build and test a complete parking domain in Go.                        |
| Library Management          | Catalog, copy vs. title, borrowing, reservations, and due dates.          | Design lending rules and test availability.                            |
| Vending Machine             | State, inventory, payment, cancellation, and change.                      | Implement valid transitions and failure recovery.                      |
| Elevator System             | Scheduling policies, requests, state, and concurrency.                    | Compare simple dispatch algorithms.                                    |
| Splitwise / Expense Sharing | Money, split strategies, balances, and rounding.                          | Implement equal, percentage, and exact splits with invariants.         |
| Notification System         | Channels, adapters, retries, and preferences.                             | Add email, SMS, and push implementations behind small interfaces.      |
| Rate Limiter                | Policy, time, concurrency, and storage trade-offs.                        | Implement a token bucket or sliding window and test concurrent access. |
| Order / Inventory / Payment | Transactions, idempotency, orchestration, and events.                     | Build a cohesive checkout workflow with failure handling.              |

## 4. Suggested 16-week pacing plan

This is a flexible plan, not a deadline. Each week should include concept
study, one implementation task, tests, a review/refactor pass, and a short
written explanation. If a task is not passing review, carry it forward rather
than rushing.

| Week | Focus                                                                   | Deliverable                                          |
| ---- | ----------------------------------------------------------------------- | ---------------------------------------------------- |
| 1    | Requirements, object modeling, encapsulation, relationships, and UML    | Parking Lot domain model and class/sequence diagrams |
| 2    | Go interfaces, composition, polymorphism, and responsibility assignment | Vehicle/Spot/Floor/Lot implementation                |
| 3    | SRP and cohesion/coupling                                               | Refactor a god object; add tests                     |
| 4    | OCP, LSP, ISP, and DIP                                                  | Policy extensibility and interface refactoring       |
| 5    | Dependency injection, packages, service layer, and repository           | Park/Exit service with fake dependencies             |
| 6    | Factory Method, Builder, and Singleton trade-offs                       | Provider/ticket creation scenario                    |
| 7    | Adapter, Decorator, Facade, and Proxy                                   | External payment adapter plus logging decorator      |
| 8    | Strategy, Observer, and State                                           | Allocation strategies and order/ticket state machine |
| 9    | Command, Chain of Responsibility, and remaining patterns                | Request pipeline and pattern-selection write-up      |
| 10   | Errors, validation, invariants, and API boundaries                      | HTTP API with validation and meaningful errors       |
| 11   | Concurrency, mutexes, race testing, and cancellation                    | Concurrent reservation tests using `go test -race`   |
| 12   | Transactions, idempotency, retries, and consistency                     | Failure-safe parking or checkout workflow            |
| 13   | Testing, fakes, integration tests, and observability                    | Test suite and instrumentation pass                  |
| 14   | Parking Lot end-to-end project                                          | Reviewed implementation, diagrams, and README        |
| 15   | Vending Machine or Library Management                                   | Second design with state/policy trade-offs           |
| 16   | Rate Limiter or Order/Inventory/Payment capstone                        | Design review, tests, and retrospective              |

## 5. Standard task template

Every coding task should include:

- **Scenario and requirements:** What the system must do, including edge cases.
- **Constraints:** Performance, concurrency, extensibility, security, or
  simplicity requirements.
- **Starter code or design prompt:** Enough context to begin without giving
  away the solution.
- **Acceptance criteria:** Observable behavior and required tests.
- **Your submission:** Proposed types and method signatures first, then
  implementation and tests.
- **Review rubric:** Correctness, responsibilities, coupling/cohesion, Go
  idioms, error handling, concurrency, test quality, and simplicity.
- **Revision:** Fix review findings and explain the trade-offs.
- **Completion gate:** Code and tests pass, design rationale is clear, and
  remaining limitations are documented.

## 6. Code review rubric

| Dimension         | Review question                                                   |
| ----------------- | ----------------------------------------------------------------- |
| Correctness       | Does it satisfy requirements and edge cases?                      |
| Responsibility    | Does each type have a coherent role and reason to change?         |
| Abstraction       | Does each interface or pattern solve a demonstrated need?         |
| Coupling/cohesion | Are dependencies and package boundaries sensible?                 |
| Go idioms         | Are interfaces small, errors explicit, and ownership clear?       |
| Concurrency       | Are shared state and race conditions handled where needed?        |
| Testing           | Are success, failure, edge, and relevant concurrent paths tested? |
| Maintainability   | Can requirements change without unrelated code changes?           |
| Simplicity        | Is complexity justified, with no speculative abstraction?         |

## 7. Progress tracker

Status legend: **Not started | In progress | Needs revision | Complete**.
Update this tracker after each learning session.

| Area                | Status             | Notes / next evidence                                                                                                   |
| ------------------- | ------------------ | ----------------------------------------------------------------------------------------------------------------------- |
| LLD foundations     | In progress        | Initial Parking Lot modeling discussion; responsibilities, encapsulation, composition, and policy separation discussed. |
| SRP                 | Next coding task   | Refactor ParkingLot responsibilities into cohesive components.                                                          |
| OCP                 | Not started        | Add a new policy without modifying stable service logic.                                                                |
| LSP                 | Not started        | Test substitutability and repair a broken abstraction.                                                                  |
| ISP                 | Not started        | Split a fat interface based on client needs.                                                                            |
| DIP / DI            | Concept introduced | `ParkingService` receives `ParkingLot` and `ParkingPolicy`; practice with code and tests.                               |
| Design patterns     | Not started        | Study each pattern through a scenario and implementation task.                                                          |
| Production concerns | Not started        | Concurrency, errors, idempotency, tests, and observability.                                                             |
| Projects            | Not started        | Parking Lot first, then varied systems.                                                                                 |

## 8. Session log template

Copy this section after each session:

- **Date / session:**
- **Topic and concept:**
- **Task attempted:**
- **What I implemented:**
- **Review findings:**
- **Tests run and results:**
- **What I learned / trade-offs:**
- **Remaining issues:**
- **Next session:**

## 9. Definition of done

A topic is complete when:

- [ ] I can explain the problem the concept solves in my own words.
- [ ] I can explain when to use it, when not to use it, and at least one
  trade-off.
- [ ] The implementation meets acceptance criteria and has meaningful tests.
- [ ] Review feedback has been addressed or explicitly documented.
- [ ] I can identify what would change if a new requirement is added.
- [ ] No pattern or abstraction was added solely because it appears in a
  syllabus.

## 10. How to resume in any study chat

Use this handoff:

> Continue my Go LLD mentorship from the syllabus artifact. Teach through
> first principles, give me one coding task per concept, do not reveal the
> complete solution upfront, review my code production-style, ask me to fix
> issues, and update the progress tracker. Current next task: SRP refactoring
> for the Parking Lot design.

**Guiding principle:** Problem -> Constraints -> Design choice -> Trade-offs ->
Code -> Tests -> Review
