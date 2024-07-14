 # Ascii_Art_Web

 - ### **Description**

    *Ascii Art Web is a web application that allows users to generate ASCII art with different fonts by entering text input. It provides a user-friendly interface for creating visually engaging text-based designs, suitable for various uses from social media posts to personal projects.*

- ### **Authors**

  -  [@oanass](https://learn.zone01oujda.ma/git/oanass)
  -  [@hsalihee](https://learn.zone01oujda.ma/git/hsalihee)

- ### **Usage**

  - #### how to run

    *Enter text into the designated area, choose your preferred font style, and then click 'Generate' to create your ASCII art.*

  - #### example

         _    _          _   _                __          __                 _       _  
        | |  | |        | | | |               \ \        / /                | |     | | 
        | |__| |   ___  | | | |   ___          \ \  /\  / /    ___    _ __  | |   __| | 
        |  __  |  / _ \ | | | |  / _ \          \ \/  \/ /    / _ \  | '__| | |  / _` | 
        | |  | | |  __/ | | | | | (_) |          \  /\  /    | (_) | | |    | | | (_| | 
        |_|  |_|  \___| |_| |_|  \___/            \/  \/      \___/  |_|    |_|  \__,_| 


- ### **Implementation details**

  - ### algorithm

    - **Input Handling:** *Begin by receiving user input, which includes the text to be converted into ASCII art and the selection of a banner.*

    - **Banner Selection:** *Identify the chosen banner from the available options. Each banner comprises a unique set of ASCII characters that will represent the input text.*

    - **Mapping Setup:** *Store the ASCII art characters (banner characters) of the selected banner in a mapping structure, typically a dictionary. Each ASCII code serves as a key, and its corresponding value is the ASCII art representation. These characters are organized in ASCII table order.*

    - **Input Processing:** *For each character in the input text:*

        *For each character in the input text Determine its ASCII code using standard encoding methods. Retrieve the corresponding ASCII art representation from the stored mapping based on the character's ASCII code. Rendering Construct the ASCII art output by concatenating the ASCII art representations retrieved for each character in the input text.*

     - **Output:** *Display or return the generated ASCII art to the user.*
