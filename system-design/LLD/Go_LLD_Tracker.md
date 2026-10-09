# Go LLD Learning Tracker

> Use this tracker alongside `Go_LLD_Syllabus_and_Study_Plan.md`.
>
> **How to mark progress:** Change `- [ ]` to `- [x]` in a Markdown
> editor. This file is an editable checklist, not a live web app;
> checkbox changes are saved when you save the file.

## Status key

- [ ] Not started
- [X] Complete

- For a topic in progress, add `**Status:** In progress` under that
  topic.
- For review feedback, add `**Needs revision:**` and list the
  remaining fixes.

## Overall dashboard

- [ ] Phase 1 --- LLD foundations
- [ ] Phase 2 --- Design principles / SOLID
- [ ] Phase 3 --- Dependency and application structure
- [ ] Phase 4 --- Creational patterns
- [ ] Phase 5 --- Structural patterns
- [ ] Phase 6 --- Behavioral patterns
- [ ] Phase 7 --- Production concerns
- [ ] Phase 8 --- Backend application architecture
- [ ] Phase 9 --- End-to-end LLD projects

**Current phase:** Phase 1 --- LLD foundations
**Current lesson/task:** Parking Lot domain model; next coding task is
SRP refactoring
**Last updated:** YYYY-MM-DD

---

## Phase 1 --- LLD foundations

- [ ] Problem decomposition and requirements
  - [ ] Write functional requirements
  - [ ] Identify constraints and edge cases
  - [ ] Define acceptance criteria
- [ ] Objects, state, behavior, and identity
  - [ ] Model domain entities
  - [ ] Assign state and behavior
- [ ] Abstraction and encapsulation
  - [ ] Protect internal state
  - [ ] Expose meaningful methods
- [ ] Relationships
  - [ ] Association
  - [ ] Aggregation
  - [ ] Composition
  - [ ] Inheritance
  - [ ] Dependency
- [ ] Polymorphism and Go interfaces
  - [ ] Define small behavior-based interfaces
  - [ ] Use interface values where variation is useful
- [ ] Composition vs inheritance
  - [ ] Identify IS-A and HAS-A relationships
  - [ ] Avoid subclass combinations for independent features
- [ ] UML and interaction flows
  - [ ] Class diagram
  - [ ] Sequence diagram
  - [ ] State diagram

## **Phase 1 notes / review findings:**

## Phase 2 --- Design principles

- [ ] SRP --- Single Responsibility Principle
  - [ ] Explain responsibility and reasons to change
  - [ ] Refactor the ParkingLot god object
  - [ ] Add tests for separated responsibilities
- [ ] OCP --- Open/Closed Principle
  - [ ] Add a new parking allocation policy
  - [ ] Avoid modifying stable service logic
- [ ] LSP --- Liskov Substitution Principle
  - [ ] Identify a broken substitutability example
  - [ ] Refactor and test implementations
- [ ] ISP --- Interface Segregation Principle
  - [ ] Identify an oversized interface
  - [ ] Split interfaces by client needs
- [ ] DIP --- Dependency Inversion Principle
  - [ ] Separate high-level policy from low-level details
  - [ ] Inject dependencies
- [ ] Coupling and cohesion
  - [ ] Review package/type responsibilities
- [ ] DRY, KISS, YAGNI
  - [ ] Compare simple and overengineered designs

## **Phase 2 notes / review findings:**

## Phase 3 --- Dependency and application structure

- [ ] Dependency Injection
- [ ] Go package design and exported APIs
- [ ] Service Layer
- [ ] Repository pattern
- [ ] Domain models vs DTOs
- [ ] Error handling and error wrapping

## **Phase 3 notes / review findings:**

## Phase 4 --- Creational patterns

- [ ] Factory Method
- [ ] Abstract Factory
- [ ] Builder
- [ ] Singleton --- including trade-offs and alternatives
- [ ] Prototype --- know when it is useful

## **Phase 4 notes / review findings:**

## Phase 5 --- Structural patterns

- [ ] Adapter
- [ ] Decorator
- [ ] Facade
- [ ] Proxy
- [ ] Composite
- [ ] Bridge --- scenario and trade-offs
- [ ] Flyweight --- scenario and trade-offs

## **Phase 5 notes / review findings:**

## Phase 6 --- Behavioral patterns

- [ ] Strategy
- [ ] Observer / pub-sub
- [ ] State
- [ ] Command
- [ ] Chain of Responsibility
- [ ] Template Method --- compare with composition in Go
- [ ] Iterator --- scenario and trade-offs
- [ ] Mediator --- scenario and trade-offs
- [ ] Memento --- scenario and trade-offs
- [ ] Visitor --- scenario and trade-offs
- [ ] Interpreter --- scenario and trade-offs

## **Phase 6 notes / review findings:**

## Phase 7 --- Production concerns

- [ ] Domain invariants and valid state transitions
- [ ] Concurrency and race conditions
  - [ ] Mutexes / atomic operations
  - [ ] Concurrent reservation test
  - [ ] Run `go test -race ./...`
- [ ] Consistency and transaction boundaries
- [ ] Retries, idempotency, and timeouts
- [ ] Logging, metrics, and observability
- [ ] Unit, integration, and contract testing
- [ ] Refactoring and production-style code reviews

## **Phase 7 notes / review findings:**

## Phase 8 --- Backend application architecture

- [ ] HTTP handlers and application APIs
- [ ] Validation and authorization
- [ ] Persistence adapters
- [ ] External service integrations
- [ ] Configuration and lifecycle management
- [ ] Modular monolith and package boundaries

## **Phase 8 notes / review findings:**

## Phase 9 --- End-to-end LLD projects

- [ ] Parking Lot
  - [ ] Requirements and class diagram
  - [ ] Domain types and responsibilities
  - [ ] Allocation policy
  - [ ] Safe reservation and release
  - [ ] Ticket lifecycle
  - [ ] Pricing and payment boundary
  - [ ] Unit and concurrency tests
  - [ ] README with design decisions and trade-offs
- [ ] Library Management System
- [ ] Vending Machine
- [ ] Elevator System
- [ ] Splitwise / Expense Sharing
- [ ] Notification System
- [ ] Rate Limiter
- [ ] Order / Inventory / Payment workflow

## **Phase 9 notes / review findings:**

---

## 16-week suggested plan

Mark a week complete only when its deliverable has been reviewed.

- [ ] **Week 1:** Requirements, object modeling, encapsulation,
  relationships, UML
- [ ] **Week 2:** Go interfaces, composition, polymorphism,
  responsibility assignment
- [ ] **Week 3:** SRP, cohesion, coupling
- [ ] **Week 4:** OCP, LSP, ISP, DIP
- [ ] **Week 5:** DI, package design, Service Layer, Repository
- [ ] **Week 6:** Factory Method, Builder, Singleton trade-offs
- [ ] **Week 7:** Adapter, Decorator, Facade, Proxy
- [ ] **Week 8:** Strategy, Observer, State
- [ ] **Week 9:** Command, Chain of Responsibility, remaining pattern
  scenarios
- [ ] **Week 10:** Errors, validation, invariants, API boundaries
- [ ] **Week 11:** Concurrency, mutexes, race tests, cancellation
- [ ] **Week 12:** Transactions, idempotency, retries, consistency
- [ ] **Week 13:** Testing, fakes, integration tests, observability
- [ ] **Week 14:** Parking Lot end-to-end project
- [ ] **Week 15:** Vending Machine or Library Management
- [ ] **Week 16:** Rate Limiter or Order/Inventory/Payment capstone

---

## Coding task log

Copy this block for each task.

### Task: \[name\]

- **Concept:**
- **Date started:**
- **Status:** Not started / In progress / Needs revision / Complete
- **Scenario and requirements:**
- **Design choice and rationale:**
- **Files / implementation:**
- **Tests written:**
- **Test command and result:**
- **Review feedback:**
- **Fixes completed:**
- **Trade-offs learned:**
- **Remaining issues:**
- **Next step:**

---

## Definition of done

- [ ] I can explain the problem this concept solves in my own words.
- [ ] I can explain when to use it and when not to use it.
- [ ] I can describe the main trade-offs.
- [ ] The code meets the acceptance criteria.
- [ ] Meaningful success, failure, and edge cases are tested.
- [ ] Review feedback is addressed or documented.
- [ ] I can explain how the design would change for a new requirement.
- [ ] I avoided unnecessary abstractions and pattern use.

## Resume prompt

> Continue my Go LLD mentorship using this tracker and the syllabus.
> Current task: SRP refactoring for the Parking Lot design. Give me one
> coding task at a time, don't reveal the complete solution upfront,
> review my Go code production-style, ask me to fix issues, and help me
> update this tracker after each reviewed task.
