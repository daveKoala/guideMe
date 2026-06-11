
## Key Design Decisions
1. **CSS-only layouts**: No Tailwind/Bootstrap - pure CSS + variables for simplicity
2. **Router metadata**: Routes specify templates via meta.template for flexibility
3. **Global reset**: Tailwind-like preflight ensures consistent baseline across browsers
4. **Named slots**: BasicLayout uses named slots for future multi-region layouts