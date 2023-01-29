# AvaloniaILSpy utils/workarounds

 - ## modules_separator
   - Problem (seems like based on [this issue](https://github.com/icsharpcode/AvaloniaILSpy/issues/66)):
     - all C# classes saving to the single file named `-Module-.cs`, instead of the supposed multiple individual `*.cs` files
   - Solution:
     - `~/some_dir$ $PATHTO/modules_separator $PATHTO/-Module-.cs`
     - produce `~/some_dir/cs-modules` directory with all classes stored separately
       - written on Go pure channels, so should be fast enough 😊
   - Installation:
     - get the latest linux x86_64 binary
     - or build it yourself using `cd ilspy_utils/cmd/modules_separator/main && go build -o modules_separator`

___

_If you want to support me - buy me a coffee via [PayPal](https://www.paypal.com/donate/?hosted_button_id=UWLQFT8HPJKK8)_