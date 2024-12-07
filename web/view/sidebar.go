package view

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// Menu item structure
type MenuItem struct {
	Label    string
	Icon     string
	SubItems []MenuItem
}

// Menu configuration
var menuItems = []struct {
	Section string
	Items   []MenuItem
}{
	{
		Items: []MenuItem{
			{Label: "Home", Icon: "home"},
			{Label: "Parking Lots", Icon: "parking"},
			{Label: "Bookings", Icon: "calendar"},
			{Label: "Parker Accounts", Icon: "users"},
			{Label: "Entries & Exits", Icon: "entries"},
			{Label: "Users", Icon: "user"},
			{Label: "RFID Tags", Icon: "rfid"},
			{Label: "Settings", Icon: "settings", SubItems: []MenuItem{
				{Label: "General", Icon: "general"},
				{Label: "Security", Icon: "security"},
				{Label: "Notifications", Icon: "notifications"},
				{Label: "Billing", Icon: "billing"},
			}},
		},
	},
}

func Sidebar() g.Node {
	return Div(
		Div(
			Class("py-2 fixed flex flex-col h-screen bg-gray-800 text-white transition-all duration-300 ease-in-out w-64 overflow-y-auto overflow-x-hidden"),
			g.Attr("x-data", `{
				mini: false,
				toggleSidebar() {
					this.mini = !this.mini;
					$el.style.width = this.mini ? '64px' : '256px';
					document.querySelector('.main-content').style.marginLeft = this.mini ? '64px' : '256px';

					// Handle collapsible elements
					$el.querySelectorAll('[x-bind\\:data-mini]').forEach(el => {
						if (this.mini) {
							el.style.width = '0';
							el.style.opacity = '0';
							el.style.visibility = 'hidden';
						} else {
							el.style.width = 'auto';
							el.style.opacity = '1';
							el.style.visibility = 'visible';
						}
					});

					// Close all submenus when collapsing
					if (this.mini) {
						$el.querySelectorAll('.submenu').forEach(submenu => {
							submenu.style.maxHeight = '0px';
							const chevron = submenu.closest('.group').querySelector('.chevron svg');
							if (chevron) chevron.style.transform = 'rotate(0deg)';
						});
					}
				},
			}`), g.Attr("x-bind:data-mini", "mini"),
			// Header section
			Div(Class("px-2"),
				// Logo and title row
				Div(Class("flex items-center"),
					// Logo container with fixed width for centering
					Div(Class("w-16 h-16 flex items-center justify-center"),
						Div(Class("w-10 h-10 bg-blue-600 rounded-lg flex items-center justify-center"),
							g.Raw(`<svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20"><path d="M2 4.75A2.75 2.75 0 014.75 2h10.5A2.75 2.75 0 0118 4.75v10.5A2.75 2.75 0 0115.25 18H4.75A2.75 2.75 0 012 15.25V4.75z"/></svg>`),
						),
					),
					// Title container
					Div(Class("flex-1 transition-all duration-300 overflow-hidden"),
						g.Attr("x-bind:data-mini", "mini"),
						Div(Class("text-lg font-semibold whitespace-nowrap"), g.Text("Cubework")),
						Div(Class("text-sm text-gray-400 whitespace-nowrap"), g.Text("Parking Admin")),
					),
					// Collapse button for expanded state
					Div(Class("flex items-center"),
						g.Attr("data-expanded-btn", "true"),
						g.Attr("x-show", "!mini"),
						g.Attr("x-transition"),
						Div(Class("cursor-pointer px-2 hover:bg-gray-700 rounded-lg"),
							g.Attr("@click", "toggleSidebar()"),
							g.Raw(`<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 19l-7-7 7-7m8 14l-7-7 7-7"/></svg>`),
						),
					),
				),
				// Collapse button for mini state
				Div(
					g.Attr("data-mini-btn", "true"),
					g.Attr("x-show", "mini"),
					g.Attr("x-cloak", ""),
					g.Attr("x-transition"),
					// Added flex container for centering
					Div(Class("flex justify-center"),
						// Fixed width container for consistent spacing
						Div(Class("w-16 h-16 flex items-center justify-center"),
							Div(Class("cursor-pointer p-2 hover:bg-gray-700 rounded-lg"),
								g.Attr("@click", "toggleSidebar()"),
								g.Raw(`<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5l7 7-7 7M5 5l7 7-7 7"/></svg>`),
							),
						),
					),
				),
			),
			// Dynamic menu sections
			g.Group(g.Map(menuItems, func(section struct {
				Section string
				Items   []MenuItem
			}) g.Node {
				return Div(Class("px-3"),
					// Section header
					Div(Class("text-sm font-medium text-gray-400 mb-2 transition-opacity duration-300"),
						g.Attr("x-bind:data-mini", "mini"),
						g.Text(section.Section),
					),
					// Section items
					g.Group(g.Map(section.Items, func(item MenuItem) g.Node {
						return NavigationItem(item)
					})),
				)
			})),

			// Footer with user profile
			Div(Class("mt-auto"),
				// Projects label - hidden in mini variant
				Div(Class("text-sm font-medium text-gray-400 px-3 mb-2 transition-all duration-300 overflow-hidden"),
					g.Attr("x-bind:data-mini", "mini"),
				),
				// User profile container
				Div(Class("flex items-center cursor-pointer hover:bg-gray-700/50 rounded-lg"),
					// Avatar container with fixed width for centering
					Div(Class("w-16 flex items-center justify-center"),
						Img(
							Class("w-8 h-8 rounded-lg"),
							Src("https://avatars.githubusercontent.com/u/124599"),
							Alt("User avatar"),
						),
					),
					// User info and dropdown - hidden in mini variant
					Div(Class("flex-1 flex items-center transition-all duration-300 overflow-hidden"),
						g.Attr("x-bind:data-mini", "mini"),
						// User info
						Div(Class("flex flex-col"),
							Div(Class("text-sm whitespace-nowrap"), g.Text("shadcn")),
							Div(Class("text-xs text-gray-400 whitespace-nowrap"), g.Text("m@example.com")),
						),
						// Dropdown arrow
						Div(Class("ml-auto pr-4"),
							g.Raw(`<svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/></svg>`),
						),
					),
				),
			),
		),
	)
}

func NavigationItem(item MenuItem) g.Node {
	hasSubmenu := len(item.SubItems) > 0
	return Div(
		Class("group relative"),
		g.Attr("x-data", `{
			expanded: false,
			toggle() {
				if(mini) return;
				if (!mini) {
					this.expanded = !this.expanded;
				}
			}
		}`),
		// Menu item button
		Div(Class("flex items-center hover:bg-gray-700/50 rounded-lg cursor-pointer px-2 py-2"),
			g.Attr("@click", "toggle"),
			// Icon container with fixed width for centering
			Div(Class("w-16 flex items-center justify-center"),
				Div(Class("w-5 h-5 text-gray-400"),
					getIcon(item.Icon),
				),
			),
			// Label and chevron container
			Div(Class("flex-1 flex items-center transition-all duration-300 overflow-hidden"),
				g.Attr("x-bind:data-mini", "mini"),
				// Label
				Div(Class("text-sm font-medium text-gray-200 whitespace-nowrap"),
					g.Text(item.Label),
				),
				// Chevron
				g.If(hasSubmenu,
					Div(Class("chevron ml-auto pr-2"),
						g.Attr("x-bind:style", `expanded ? 'transform: rotate(90deg)' : ''`),
						g.Raw(`<svg class="w-4 h-4 text-gray-400 transition-transform duration-200" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>`),
					),
				),
			),
		),
		// Submenu
		g.If(hasSubmenu,
			Div(Class(`
				submenu overflow-hidden transition-all duration-200
				data-[mini="false"]:max-h-0
				data-[mini="true"]:hidden group-hover:data-[mini="true"]:block
			`),
				g.Attr("x-cloak", ""),
				g.Attr("x-bind:style", `expanded ? 'max-height: ' + $el.scrollHeight + 'px' : 'max-height: 0px'`),
				Div(Class("pl-10 py-1 space-y-1 data-[mini='true']:pl-0"),
					g.Group(g.Map(item.SubItems, func(subItem MenuItem) g.Node {
						return Div(Class("flex items-center gap-3 px-3 py-2 hover:bg-gray-700/50 rounded-lg cursor-pointer"),
							Div(Class("w-5 h-5 text-gray-400"),
								getIcon(subItem.Icon),
							),
							Div(Class("text-sm font-medium text-gray-200"),
								g.Text(subItem.Label),
							),
						)
					})),
				),
			),
		),
	)
}

// Helper function to get icons (you can add more icons as needed)
func getIcon(name string) g.Node {
	icons := map[string]string{
		"home":          `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/></svg>`,
		"parking":       `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5h14a2 2 0 012 2v10a2 2 0 01-2 2H5a2 2 0 01-2-2V7a2 2 0 012-2zm0 0l5 5m9-5l-5 5m5 0l-5 5m-9-5l5 5"/></svg>`,
		"calendar":      `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>`,
		"users":         `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/></svg>`,
		"entries":       `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4"/></svg>`,
		"user":          `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>`,
		"rfid":          `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z"/></svg>`,
		"settings":      `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/></svg>`,
		"general":       `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"/></svg>`,
		"security":      `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/></svg>`,
		"notifications": `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/></svg>`,
		"billing":       `<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"/></svg>`,
	}
	return g.Raw(icons[name])
}
