# [heavens-vault-vault](https://github.com/ddbeck/heavens-vault-vault)

by [@ddbeck](https://www.ddbeck.com/)

`heavens-vault-vault` makes a copy of your [_Heaven's Vault_](https://www.inklestudios.com/heavensvault/) save game every ten minutes, so you can revert your game after you've discovered that you've made some unrecoverable mistake and you don't want to start a whole new playthrough.

I created this program to ease some frustration while playing and to learn a little bit about the Go programming language. I decided to share it in case it was useful to anyone else, but I don't expect to add any features or futher develop this program. If you have any problems or ideas, feel free to open an issue or pull request, but I can't promise I'll be able to follow up on them.

You're free to use, copy, modify, and distribute `heavens-vault-vault` under the terms of the GPL v3. See `LICENSE` for details.

## Build

**Note**: you can probably skip this step. Download the executable from [the releases page](https://github.com/ddbeck/heavens-vault-vault/releases/) instead, then mosey on down to [Setup](#setup).

To build `heavens-vault-vault`, you must have [Go](https://golang.org/) installed (I used Go 1.12—other versions may work, but I'm not sure because I'm new to Go).

1. Clone this repository.

2. In the root of the repository, run `go build -o heavens-vault-vault.exe`.

   If you want to cross-compile from another platform (I wrote the program and compiled it on a Mac), then run `GOOS=windows GOARCH=amd64 go build -o heavens-vault-vault.exe` and copy the executable to the computer where you're playing _Heaven's Vault_.


## Setup

To set up the program:

1. In _File Explorer_, right click `heavens-vault-vault.exe`, then click **Create shortcut**.

2. Right-click on the new shortcut, then click **Properties**.

3. In the _Target_ field, put quotes (`"`) around the complete path to `heavens-vault-vault.exe`, then add a space, and the path to the foler where your `heavensVaultSave.json` file is, also in quotes.
 
   For example, my Heaven's Vault save game is in `C:/Users/ddbeck/AppData/LocalLow/Inkle Ltd/Heaven's Vault/`; my _Target_ field contains: `"C:/Users/ddbeck/src/heavens-vault-vault/heavens-vault-vault.exe" "C:/Users/ddbeck/AppData/LocalLow/Inkle Ltd/Heaven's Vault/"`.

4. Click **OK**.

## Usage

1. Before starting _Heaven's Vault_, double click [the shortcut you created previously](#setup). A new window appears with a running status of the program.

   Every ten minutes, `heavens-vault-vault.exe` copies your current `heavensVaultSave.json` to a new folder, `hvvBackupSaves`. Each copy is named with the date and time.

2. Play _Heaven's Vault_.

   If you decide you want to go back to a previous save game, quit _Heaven's Vault_, replace `heavensVaultSave.json` with a file from `hvvBackupSaves`, and restart _Heaven's Vault_.

3. When you're done playing _Heaven's Vault_, select the window for `heavens-vault-vault.exe` and press `Ctrl + c` to quit.
