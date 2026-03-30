# FileModes -- permissions in Go

## Common Permissions
| Mode | Purpose                | Explanation                                     |
|------|------------------------|-------------------------------------------------|
| 0600 | Private file           | Owner can read/write, no access for others      |
| 0644 | Normal file            | Owner can read/write, group and others can read |
| 0666 | Shared writable file   | Everyone can read/write, no execute             |
| 0700 | Private executable/dir | Only owner has full access                      |
| 0755 | Public executable/dir  | Owner full access, others can read/execute      |
| 0777 | Open access            | Everyone can read/write/execute                 |

## Permissions defined at each level
| Octal | Meaning        | Purpose              | Explanation                                  |
|-------|----------------|----------------------|----------------------------------------------|
| 0400  | Owner read     | Owner can read       | File owner can view the file contents        |
| 0200  | Owner write    | Owner can modify     | File owner can edit or overwrite the file    |
| 0100  | Owner execute  | Owner can run        | Needed to execute a file or enter a directory|
| 0040  | Group read     | Group can read       | Users in the file's group can view it        |
| 0020  | Group write    | Group can modify     | Users in the group can change it             |
| 0010  | Group execute  | Group can run        | Group can execute file or access directory   |
| 0004  | Others read    | Everyone can read    | Any other user can view it                   |
| 0002  | Others write   | Everyone can modify  | Any other user can change it                 |
| 0001  | Others execute | Everyone can run     | Any other user can execute/access it         |
