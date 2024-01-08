const navLinks: NodeListOf<Element> = document.querySelectorAll("nav a");

for (let i : number = 0; i < navLinks.length; i++) {
    const link : Element = navLinks[i];

    if (link.getAttribute('href') === window.location.pathname) {
        link.classList.add("live");
        break;
    }
}