<h1 align="center" id="title">MagicMerge</h1>

<p id="description">MagicMerge is a sophisticated utility designed specifically for cybersecurity professionals and red team operators. It offers a powerful yet covert method for embedding and encrypting executables within images enabling secure payload delivery and execution in environments where discretion and evasion are paramount.</p>

  
  
<h2>🧐 Features</h2>

Here're some of the project's best features:

*   Stealth Integration: MagicMerge excels in concealing executables within image files making it an ideal tool for creating steganographic payloads that bypass common detection techniques used in digital forensics and incident response.
*   Robust Encryption: Incorporates a robust 16-byte XOR encryption mechanism with dynamically generated keys ensuring that embedded executables are shielded from unauthorized access and analysis. This feature is crucial for maintaining operational security and payload integrity in hostile environments.
*   Flexible Deployment Options: Whether deploying internally within a controlled lab environment or during an active red team engagement MagicMerge provides the option to serve payloads locally or via a public URL using ngrok. This flexibility allows operators to adapt to various operational scenarios and network constraints.
*   Command-Line Efficiency: Designed for rapid use in high-stakes environments MagicMerge operates entirely via command line allowing for quick adjustments scripting and integration with other tools in a cybersecurity professional’s toolkit.
*   Cross-Platform Compatibility: Built in Go MagicMerge runs seamlessly across Windows macOS and Linux ensuring reliable deployment regardless of the host operating system used during operations.

<h2>🛠️ Installation Steps:</h2>

<p>1. Make sure to set Ngrok in</p>

```
%PATH%
```

<p>2. Install go libraries</p>

```
go get github.com/fatih/color
```

```
go get github.com/logrusorgru/aurora/v3
```
