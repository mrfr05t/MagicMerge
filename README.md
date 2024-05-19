<h1 align="center" id="title">MagicMerge</h1>

<p id="description">MagicMerge is a sophisticated utility designed specifically for cybersecurity professionals and red team operators. It offers a powerful yet covert method for embedding and encrypting executables within images enabling secure payload delivery and execution in environments where discretion and evasion are paramount.</p>

  
  
<h2>🧐 Features</h2>

Here're some of the project's best features:
Stealthy Embedding
•	Description: MagicMerge allows for the embedding of executables directly within image files, creating an unobtrusive method of delivering payloads that bypasses common security measures.
•	Benefit: Increases operational success by evading detection from security systems that scan for suspicious binary formats.
Robust XOR Encryption
•	Description: MagicMerge encrypts executables embedded within images using a dynamically generated 16-byte XOR key, ensuring robust protection against unauthorized access and analysis.
•	Benefit: Maintains the integrity and confidentiality of the payload, protecting sensitive data from potential breaches.
Flexible Deployment Options
•	Description: The tool supports both local and ngrok-enabled public hosting, allowing payloads to be accessed over internal networks or the internet, depending on operational needs.
•	Benefit: Offers versatility in deployment, making it suitable for a range of scenarios from controlled internal tests to wide-reaching external deployments.
Command-Line Interface
•	Description: Operates entirely via a command-line interface, facilitating rapid deployment, scripting, and integration with existing cybersecurity tool chains.
•	Benefit: Enhances usability and efficiency, allowing users to execute operations quickly and automate processes as needed.
Cross-Platform Compatibility
•	Description: Built in Go, MagicMerge is inherently cross-platform, running seamlessly on Windows, macOS, and Linux.
•	Benefit: Ensures consistent performance and functionality across different operating systems, which is crucial for teams working in diverse IT environments.
Easy Configuration and Customization
•	Description: Simple configuration options via command-line flags enable quick setup and adjustments to fit specific operational requirements.
•	Benefit: Allows for easy customization and fine-tuning, improving operational effectiveness and adaptability.
Dynamic XOR Key Generation
•	Description: Generates a new XOR key for each session, enhancing the security of each encrypted payload.
•	Benefit: Prevents the compromise of one payload from affecting others, safeguarding multiple deployments simultaneously.
Random Image Fetching
•	Description: Automatically fetches a new, random image for each payload from a designated source to ensure each payload is uniquely camouflaged.
•	Benefit: Increases security by preventing pattern recognition in payload deliveries, crucial for operations requiring high stealth.
Flexible Payload Sourcing
•	Description: Allows users to source executables from URLs or local storage, accommodating diverse input needs and operational environments.
•	Benefit: Adapts to varying scenarios easily, from centrally managed payloads that can be updated remotely to securely handled local files for restricted environments.



<h2>🛠️ Installation Steps:</h2>

<p>1. Download and Install Go (Windows)</p>

```
https://go.dev/doc/install
```

<p>2. Install go libraries</p>

```
go get github.com/fatih/color
```

```
go get github.com/logrusorgru/aurora/v3
```

<p>4. Set ngrok in Enviorment Variable path</p>

```
%PATH% 
```
