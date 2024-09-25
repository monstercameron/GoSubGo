// expanded-htmx-like-script.js
document.addEventListener("DOMContentLoaded", function() {
    const debounceTimers = {};
    const intervalTimers = {};

    function handleEvent(event) {
        const element = event.currentTarget;
        const eventType = event.type;
        const url = element.dataset.url;
        const target = element.dataset.target;
        const swap = element.dataset.swap;
        const delay = parseInt(element.dataset.delay) || 0;
        const timeout = parseInt(element.dataset.timeout) || 0;
        const debounce = parseInt(element.dataset.debounce) || 0;
        const interval = parseInt(element.dataset.interval) || 0;

        if (element.dataset.preventDefault !== undefined) {
            event.preventDefault();
        }

        if (element.dataset.confirm) {
            if (!confirm(element.dataset.confirm)) {
                return;
            }
        }

        let params = {};
        if (element.dataset.params) {
            try {
                params = JSON.parse(element.dataset.params);
            } catch (e) {
                console.error("Invalid JSON in data-params:", e);
            }
        }

        if (element.tagName.toLowerCase() === 'form') {
            const formData = new FormData(element);
            for (const [key, value] of formData.entries()) {
                params[key] = value;
            }
        }

        const eventData = {
            eventType: eventType,
            elementID: element.id,
            url: url,
            target: target,
            swap: swap,
            params: params
        };

        const sendRequest = () => {
            if (element.dataset.disable !== undefined) {
                element.disabled = true;
            }

            const indicator = element.dataset.indicator;
            if (indicator) {
                document.querySelector(indicator).classList.remove('hidden');
            }

            const controller = new AbortController();
            const signal = controller.signal;

            if (timeout > 0) {
                setTimeout(() => controller.abort(), timeout);
            }

            if (typeof window.handleEvent === 'function') {
                window.handleEvent(JSON.stringify(eventData), signal);
            } else {
                console.error("handleEvent function is not available.");
            }
        };

        if (debounce > 0) {
            clearTimeout(debounceTimers[element.id]);
            debounceTimers[element.id] = setTimeout(sendRequest, debounce);
        } else if (interval > 0) {
            clearInterval(intervalTimers[element.id]);
            sendRequest();
            intervalTimers[element.id] = setInterval(sendRequest, interval);
        } else {
            setTimeout(sendRequest, delay);
        }
    }

    function attachEventListeners() {
        const elements = document.querySelectorAll('[data-trigger]');
        elements.forEach(element => {
            const triggers = element.dataset.trigger.split(' ');
            triggers.forEach(trigger => {
                element.addEventListener(trigger, handleEvent);
            });
        });
    }

    attachEventListeners();

    // Handle the response from WASM
    window.updateDOM = function(targetSelector, content, swapMethod, pushUrl) {
        const target = document.querySelector(targetSelector);
        if (!target) {
            console.error(`Target element not found: ${targetSelector}`);
            return;
        }

        switch (swapMethod) {
            case 'innerHTML':
                target.innerHTML = content;
                break;
            case 'outerHTML':
                target.outerHTML = content;
                break;
            case 'beforebegin':
                target.insertAdjacentHTML('beforebegin', content);
                break;
            case 'afterbegin':
                target.insertAdjacentHTML('afterbegin', content);
                break;
            case 'beforeend':
                target.insertAdjacentHTML('beforeend', content);
                break;
            case 'afterend':
                target.insertAdjacentHTML('afterend', content);
                break;
            case 'class':
                target.className = content;
                break;
            default:
                console.error(`Unknown swap method: ${swapMethod}`);
        }

        if (pushUrl) {
            history.pushState(null, '', pushUrl);
        }

        // Re-enable elements and hide indicators
        document.querySelectorAll('[data-disable]').forEach(el => el.disabled = false);
        document.querySelectorAll('[data-indicator]').forEach(el => {
            const indicator = document.querySelector(el.dataset.indicator);
            if (indicator) {
                indicator.classList.add('hidden');
            }
        });

        // Re-attach event listeners to new elements
        attachEventListeners();
    };

    // Trigger 'load' events
    document.querySelectorAll('[data-trigger="load"]').forEach(element => {
        handleEvent({ currentTarget: element, type: 'load' });
    });
});