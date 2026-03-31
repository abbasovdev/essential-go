<div align="center">
    <h1>5) Project Info</h1>
    <a class="header-badge" target="_blank" href="https://www.linkedin.com/in/abbasovdev/">
        <img alt="Linkedin Follow" src="https://img.shields.io/badge/style--5eba00.svg?label=LinkedIn&logo=linkedin&style=social">
    </a>
    <a class="header-badge" target="_blank" href="https://x.com/abbcyhn">
        <img alt="X Follow" src="https://img.shields.io/twitter/follow/abbcyhn?style=social">
    </a>
    <h2>Author: 
        <a href="https://www.linkedin.com/in/abbasovdev/" target="_blank">Jeyhun Abbasov</a>
    </h2>
    <div>
        <span>Interactive Learning</span>
        <br />
        <div align="left">
            ☰ <a href="../../../README.md">🏠 Home</a> 
            ┃ <a href="../03_how_to_use/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../02_basics/01_hello_world/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>


---


## Project Info

Besides writing code in each topic, you also will build a project.

You will build a CLI Word Analytics Tool called `Grolyze`. 

A tool where a user runs  `./grolyze -file=guide.md` and gets a structured report on word frequency, unique count, and formatting statistics.


## Project Roadmap

Here is what you will improve in each chapter to build the final tool.

⮕&#xFE0F; `Chapter 2. Basics`

* Create `main.go`
* Print "Grolyze: Guide Analyzer Initialized"


⮕&#xFE0F; `Chapter 3. Variables`

* Define constants for the tool version 
* Create variables to hold the "target file" string and "is verbose" boolean
* Use explicit and implicit typing for variables

⮕&#xFE0F; `Chapter 4. Data Structures`

* Use a Slice to store a list of "stop words" (words like "the", "is", "a")
* Use a Map to store words as keys and their frequency count as values
* Practice how Go handles dynamic collections

⮕&#xFE0F; `Chapter 5. Flow Controls`

* Use a for loop to iterate over the raw words extracted from the guide files 
* Use if statements to check if a word is in your "stop words" slice
* Use switch to categorize words by length (e.g., short, medium, long)
* Practice logical branching and range-based loops

⮕&#xFE0F; `Chapter 6. Functions`

* Create a function count words
* Use Multiple Returns to return the map and the total word count
* Use Defer to print a "Analysis Complete" message 
* Practice refactoring the code

⮕&#xFE0F; `Chapter 7. Pointers`

* Pass word count map to function as a pointer
* Practice passing by reference and value 

⮕&#xFE0F; `Chapter 8. Structs & Methods`

* Define a AnalysisReport struct
* Move the logic into Methods attached to this struct 
* Export fields, so they can be accessed from a different package
* Practice Go's OOP approach 

⮕&#xFE0F; `Chapter 9. Interfaces`

* Define a Formatter interface for output formats
* Implement TextFormatter and JSONFormatter
* Use interfaces as function parameters

⮕&#xFE0F; `Chapter 10. Generics`

* Add generic Filter and Keys utility functions
* Use type parameters and constraints
* Replace stop word filtering with generic Filter

⮕&#xFE0F; `Chapter 11. Errors`

* Implement real file reading using os.ReadFile
* Use the idiomatic if err != nil pattern to handle errors
* Practice explicit error handling

⮕&#xFE0F; `Chapter 12. Goroutines`

* Process multiple files concurrently using goroutines
* Use channels to collect results from goroutines
* Practice Go's built-in concurrency model

⮕&#xFE0F; `Chapter 13. Conclusion`

* Set up the go.mod file, pkg, and internal
* Practice how to use modules

## Conclusion

Now you know what project you will be building during this guide.

Continue to the next topic 👇

[Hello World ⮕&#xFE0F;](../../02_basics/01_hello_world/README.md)