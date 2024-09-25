# GoSubGo: An Experimental Event-Driven Architecture with Go and WebAssembly

Welcome to **GoSubGo**! 🎉 This project is an experimental exploration into building a modern, event-driven web application using Go and WebAssembly (WASM). By leveraging the strengths of Go and JavaScript, GoSubGo aims to create a seamless and efficient development experience that bridges the gap between client-side interactivity and server-side performance.

**Note:** This project is in its experimental stages. We’re excited to share our journey and welcome contributions, feedback, and discussions to help shape its future!

## Table of Contents

- [Introduction](#introduction)
- [Key Concepts](#key-concepts)
  - [HATEOAS (Hypermedia as the Engine of Application State)](#hateoas-hypermedia-as-the-engine-of-application-state)
  - [Go vs. JavaScript](#go-vs-javascript)
  - [Functional Programming](#functional-programming)
  - [Go Error Handling](#go-error-handling)
  - [Event-Driven Architecture](#event-driven-architecture)
  - [Goroutines and Performance](#goroutines-and-performance)
  - [Shared Objects Between Client and Server](#shared-objects-between-client-and-server)
  - [SQL on Frontend and Backend](#sql-on-frontend-and-backend)
- [Design Questions](#design-questions)
- [Additional Concepts](#additional-concepts)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Building the WASM Module](#building-the-wasm-module)
  - [Running the Application](#running-the-application)
- [Contributing](#contributing)
- [License](#license)

## Introduction

**GoSubGo** is an ambitious project that seeks to harness the power of Go and WebAssembly to create dynamic, high-performance web applications. By integrating Go's robust backend capabilities with the flexibility of JavaScript on the frontend, GoSubGo aims to provide a unified development experience where both client and server share common logic and data structures.

This project delves into several advanced concepts, including HATEOAS, event-driven architectures, and the utilization of SQL databases on both the frontend and backend. Through this experimentation, we hope to uncover best practices and innovative approaches to modern web development.

## Key Concepts

### HATEOAS (Hypermedia as the Engine of Application State)

HATEOAS is a constraint of the REST application architecture that keeps the RESTful style unique from other network application architectures. In GoSubGo, we explore how hypermedia can drive the interaction between the client and server, allowing the client to dynamically navigate the application's state through hyperlinks provided by the server.

### Go vs. JavaScript

**Go** and **JavaScript** are powerful languages, each with its own strengths:

- **Go:**
  - **Performance:** Compiled to native machine code, offering superior performance.
  - **Concurrency:** Goroutines provide lightweight, efficient concurrency.
  - **Strong Typing:** Enhances code reliability and maintainability.
  
- **JavaScript:**
  - **Flexibility:** Highly dynamic and versatile for frontend interactions.
  - **Ubiquity:** Native to all modern web browsers.
  - **Asynchronous Programming:** Promises and async/await simplify handling asynchronous operations.

GoSubGo leverages Go's performance and concurrency features while maintaining the flexibility and interactivity of JavaScript on the frontend through WebAssembly.

### Functional Programming

Functional programming (FP) emphasizes immutability, first-class functions, and pure functions. In GoSubGo, we incorporate FP principles to enhance code modularity, readability, and maintainability. This includes using higher-order functions, avoiding side effects, and leveraging immutable data structures where possible.

### Go Error Handling

Go’s approach to error handling is explicit and straightforward, relying on error returns rather than exceptions. This clarity helps in writing robust and predictable code. In GoSubGo, we embrace this paradigm to handle errors gracefully across both client and server components, ensuring that failures are managed effectively without unexpected crashes.

### Event-Driven Architecture

An event-driven architecture (EDA) allows the system to respond to events asynchronously, promoting scalability and flexibility. GoSubGo adopts EDA to manage interactions between the frontend and backend. Events triggered by user actions are captured by JavaScript, dispatched to the Go event bus, and processed accordingly. This decoupled approach facilitates easier maintenance and scalability.

### Goroutines and Performance

**Goroutines** are lightweight threads managed by Go’s runtime, enabling high concurrency with minimal resource overhead. GoSubGo utilizes goroutines to handle multiple tasks concurrently, such as processing user events, managing database operations, and handling network requests, thereby enhancing the application's performance and responsiveness.

### Shared Objects Between Client and Server

One of the innovative aspects of GoSubGo is the potential to share objects and data structures between the client and server. This uniformity reduces duplication, ensures consistency, and simplifies the development process by allowing both sides to operate on the same data models and logic.

### SQL on Frontend and Backend

GoSubGo employs **SQL.js** on the frontend and SQLite on the backend, enabling a unified approach to data management. This dual usage of SQL allows for powerful querying capabilities on both client and server, facilitating features like offline support, real-time data synchronization, and complex data manipulations.

## Design Questions

As we embark on this experimental journey with GoSubGo, several intriguing questions guide our development process:

1. **How should the client and server communicate effectively?**
   - What protocols and data formats will best facilitate seamless interaction between Go (WASM) and JavaScript?

2. **Can we send SQLite patches to the backend to maintain application state across logins?**
   - How can we efficiently serialize and transmit database changes from the client to the server?

3. **Is maintaining a separate SQL database per user a viable strategy to prevent SQL injections?**
   - How does this approach impact scalability and data management?

4. **How can an admin query multiple SQLite instances simultaneously as if they were a single entity?**
   - What mechanisms can aggregate and interface with multiple databases seamlessly?

5. **What are the best practices for error handling across the client-server boundary in a WebAssembly context?**
   - How can we ensure that errors are communicated clearly and handled gracefully?

6. **How can we leverage Go’s concurrency model (goroutines) to optimize performance in a web environment?**
   - What patterns can maximize the benefits of goroutines without introducing complexity?

7. **What strategies can ensure that shared objects between client and server remain consistent and synchronized?**
   - How can we manage state changes to prevent conflicts and ensure data integrity?

8. **How can SQL be effectively utilized on both the frontend and backend to provide a cohesive data management experience?**
   - What synchronization techniques are necessary to keep data consistent across both layers?

These questions not only drive the development of GoSubGo but also open avenues for deeper exploration and learning within the realms of web development, concurrency, and data management.

## Additional Concepts

In addition to the primary concepts outlined above, GoSubGo touches upon several other important areas:

- **Hygiene and Security:**
  - Ensuring secure data transmission between client and server.
  - Protecting against common web vulnerabilities.

- **Scalability:**
  - Designing the architecture to handle increasing loads gracefully.
  - Efficient resource management with goroutines and concurrency.

- **Modularity and Reusability:**
  - Structuring the codebase to promote reusable components.
  - Facilitating easier testing and maintenance.

- **User Experience (UX):**
  - Creating a responsive and intuitive interface.
  - Providing real-time feedback and seamless interactions.

- **State Management:**
  - Maintaining consistent application state across different components.
  - Handling state synchronization between client and server.

## Project Structure

Understanding the project's structure is key to navigating and contributing effectively. Here's an overview:

```
GoSubGo/
├── bin/
│   └── main.wasm            # Compiled WebAssembly module
├── scripts/
│   ├── script.js            # Custom JavaScript for event handling
│   ├── wasm_exec.js         # Go's WASM JavaScript bridge
│   └── sql-wasm.js          # SQL.js library
├── src/
│   ├── database/
│   │   └── database.go      # Database layer using SQL.js and SQLite
│   ├── events/
│   │   └── events.go        # Event bus implementation
│   ├── todolist/
│   │   └── todos.go         # Todo logic and rendering
│   └── utils/
│       └── utils.go         # Utility functions for DOM manipulation
├── index.html               # Main HTML file
├── main.go                  # Go entry point
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies
└── README.md                # Project README
```

## Getting Started

Embarking on this experimental project is both exciting and challenging. Here's how you can get started:

### Prerequisites

- **Go 1.22.0 or higher**: Required for building the WASM module.
- **Node.js and NPM**: Useful for setting up a local development server (optional).
- **Modern Web Browser**: Supports WebAssembly (e.g., Chrome, Firefox, Edge, Safari).

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/monstercameron/GoSubGo.git
   cd GoSubGo
   ```

2. **Install Dependencies:**

   - Ensure that the `sql-wasm.js` and `sql-wasm.wasm` files are available in the `scripts` directory. These can be obtained from the [SQL.js](https://github.com/sql-js/sql.js/) repository.

### Building the WASM Module

1. **Set Go Environment Variables:**

   ```bash
   export GOOS=js
   export GOARCH=wasm
   ```

2. **Build the WASM Module:**

   ```bash
   go build -o ./bin/main.wasm main.go
   ```

3. **Copy `wasm_exec.js`:**

   Copy the `wasm_exec.js` file from your Go installation to the `scripts` directory:

   ```bash
   cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./scripts/
   ```

### Running the Application

Browsers impose restrictions on accessing local files via JavaScript modules. Therefore, serving the application using a local HTTP server is recommended.

1. **Using Python's Simple HTTP Server:**

   ```bash
   python -m http.server 8080
   ```

2. **Using Node.js:**

   ```bash
   npx http-server -p 8080
   ```

3. **Access the Application:**

   Open your browser and navigate to `http://localhost:8080/` to see GoSubGo in action.

## Contributing

We’re thrilled about the potential of GoSubGo and welcome contributions from the community! Whether you’re looking to fix bugs, add new features, or simply provide feedback, your involvement is invaluable.

### How to Contribute

1. **Fork the Repository:**

   Click the "Fork" button at the top of the repository page to create your own copy.

2. **Clone Your Fork:**

   ```bash
   git clone https://github.com/yourusername/GoSubGo.git
   cd GoSubGo
   ```

3. **Create a Feature Branch:**

   ```bash
   git checkout -b feature/my-new-feature
   ```

4. **Commit Your Changes:**

   ```bash
   git commit -am 'Add new feature'
   ```

5. **Push to the Branch:**

   ```bash
   git push origin feature/my-new-feature
   ```

6. **Open a Pull Request:**

   Navigate to your forked repository on GitHub and open a pull request to the main GoSubGo repository.

### Reporting Issues

Encountered a bug or have a feature request? Please open an issue in the [GitHub Issues](https://github.com/monstercameron/GoSubGo/issues) section. Provide as much detail as possible to help us understand and address the problem effectively.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute it as per the terms of the license.

---

Embarking on the GoSubGo journey is both humbling and exhilarating. As we experiment with cutting-edge technologies and architectural patterns, we invite you to join us in exploring the possibilities of Go and WebAssembly. Together, we can push the boundaries of what's achievable in web development. Happy coding! 🚀