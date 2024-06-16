# Go_Reloaded
This project is a text modification tool written in Go, designed to perform various text transformations based on specific markers. The tool takes an input file containing text and applies the modifications, saving the result to an output file. The code follows Go best practices and includes unit tests to ensure functionality.

## Features
The tool performs the following modifications:

### Hexadecimal to Decimal Conversion

#### Converts words followed by (hex) to their decimal equivalents.
Example: 1E (hex) files were added -> 30 files were added
Binary to Decimal Conversion

#### Converts words followed by (bin) to their decimal equivalents.
Example: It has been 10 (bin) years -> It has been 2 years
Uppercase Conversion

#### Converts words followed by (up) to uppercase.
Example: Ready, set, go (up) ! -> Ready, set, GO !
Lowercase Conversion

#### Converts words followed by (low) to lowercase.
Example: I should stop SHOUTING (low) -> I should stop shouting
Capitalization

#### Capitalizes words followed by (cap).
Example: Welcome to the Brooklyn bridge (cap) -> Welcome to the Brooklyn Bridge
Multiple Word Conversion

#### Applies uppercase, lowercase, or capitalization to a specified number of preceding words.
Example: This is so exciting (up, 2) -> This is SO EXCITING
Punctuation Formatting

#### Ensures punctuation marks are correctly spaced.
Example: I was sitting over there ,and then BAMM !! -> I was sitting over there, and then BAMM!!
Ellipsis and Combined Punctuation

#### Properly formats ellipses and combined punctuation.
Example: I was thinking ... You were right -> I was thinking... You were right
A/An Correction

#### Corrects the usage of 'a' and 'an' based on the following word.
Example: There it was. A amazing rock! -> There it was. An amazing rock!
Apostrophe Handling

#### Places apostrophes correctly around words.
Example: I am exactly how they describe me: ' awesome ' -> I am exactly how they describe me: 'awesome'
Example: As Elton John said: ' I am the most well-known homosexual in the world ' -> As Elton John said: 'I am the most well-known homosexual in the world'
Usage

To run the tool and apply modifications to your text:

Basic Usage :

```sh
go run . input.txt output.txt
```
