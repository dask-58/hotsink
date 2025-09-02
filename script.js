(function() {
    const root = document.documentElement;
    const storageKey = 'theme-preference';

    function getPreferredTheme() {
        const stored = localStorage.getItem(storageKey);
        if (stored) return stored;
        return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    }

    function setTheme(theme) {
        if (theme === 'dark') {
            root.setAttribute('data-theme', 'dark');
        } else {
            root.removeAttribute('data-theme');
        }
    }

    setTheme(getPreferredTheme());

    const btn = document.getElementById('themeToggle');
    if (btn) {
        btn.addEventListener('click', function() {
            const next = root.getAttribute('data-theme') === 'dark' ? 'light' : 'dark';
            setTheme(next);
            localStorage.setItem(storageKey, next);
        });
    }

    function toggleLog(headerElement) {
        const logEntry = headerElement.closest('.log-entry');
        logEntry.classList.toggle('collapsed');
    }

    window.toggleLog = toggleLog;

    document.addEventListener('DOMContentLoaded', function() {
        const logEntries = document.querySelectorAll('.log-entry.collapsible');
        logEntries.forEach(entry => {
            entry.classList.add('collapsed');
        });
    });
})();

