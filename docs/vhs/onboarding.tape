Output docs/assets/onboarding_demo.gif
Set FontSize 24
Set Width 1000
Set Height 800
Set Padding 30
Set Theme "TokyoNight"

# Warm up
Sleep 1s

# Launch
Type "gyanesh-help onboarding"
Enter
Sleep 3s

# First prompt: "Press [Enter] or type 'y' to continue..."
# We use Enter here.
Enter
Sleep 3s

# Second prompt: "Would you like to install a stack right now?"
# Type y to trigger the store picker
Type "y"
Sleep 500ms
Enter
Sleep 3s

# TUI interaction
Down
Sleep 800ms
Down
Sleep 800ms
Up
Sleep 800ms
Type "q"
Sleep 2s

# Final sequence
Sleep 10s
