# AvaloniaILSpy utils/workarounds

 - ## modules_separator
   - Problem (seems like based on [this issue](https://github.com/icsharpcode/AvaloniaILSpy/issues/66)):
     - all C# classes saving to the single file named `-Module-.cs`, instead of the supposed multiple individual `*.cs` files
   - Solution:
     1. **_recommended_**: based on [nemec comment](https://github.com/icsharpcode/AvaloniaILSpy/issues/66#issuecomment-1032213132) you could get separated `*cs` files if click on root "Assembly-CSharp ..." then "Save code ..." where you have to add `.csproj` so it should be `Assembly-CSharp.csproj` in the first text field && selected `C# project`
     2. use `modules_separator` to split each generated codebase (in ILSpy you can't save more than one selected node, so this way actual if you want save only particular subroot node)
         - `~/some_dir$ $PATHTO/modules_separator $PATHTO/-Module-.cs`
         - produce `~/some_dir/cs-modules` directory with all classes stored separately
             - written on Go pure channels, so should be fast enough 😊
   - Installation:
     - [get the latest linux x86_64 binary](https://github.com/madzohan/AvaloniaILSpy-utils/releases/download/1.0.0/modules_separator)
     - or build it yourself using `cd ilspy_utils/cmd/modules_separator/main && go build -ldflags "-w" -o modules_separator`

___

_If you want to support me - buy me a coffee via [PayPal](https://www.paypal.com/donate/?hosted_button_id=UWLQFT8HPJKK8)_
