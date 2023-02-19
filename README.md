# Koala Violet 
## PNG Image Compressor

Tool to compress png images
This is a Go program that compresses PNG images. The program uses the `image/png` package to load and compress PNG images.

## Usage

place the png images to compress in the input directory, they can be in nested folders

``This will look for all PNG files in the input directory and its subdirectories, and create a resized copy of each file in a directory called `output` in the same directory as the executable.
``
The tool uses the `MitchellNetravali` interpolation algorithm to resize the images. You can modify this by editing the code in `main.go`.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please create an issue on GitHub. If you want to contribute code, please fork the repository and create a pull request.

## License

This tool is licensed under the MIT License. See the `LICENSE` file for more information.