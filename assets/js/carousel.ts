// Carousel Logic
document.addEventListener('DOMContentLoaded', () => {
    const carouselContainer = document.querySelector('[data-carousel]');
    if (!carouselContainer) return;

    const slides = carouselContainer.querySelectorAll('[data-slide]');
    const indicators = carouselContainer.querySelectorAll('[data-indicator]');
    const prevBtn = carouselContainer.querySelector('[data-prev]');
    const nextBtn = carouselContainer.querySelector('[data-next]');
    
    let currentIndex = 0;
    const totalSlides = slides.length;

    function showSlide(index) {
        // Handle wrapping
        if (index < 0) {
            currentIndex = totalSlides - 1;
        } else if (index >= totalSlides) {
            currentIndex = 0;
        } else {
            currentIndex = index;
        }

        // Update slides visibility
        slides.forEach((slide, i) => {
            if (i === currentIndex) {
                slide.classList.remove('hidden');
                slide.classList.add('block');
                slide.classList.add('animate-fade-in'); // Add fade-in animation
            } else {
                slide.classList.add('hidden');
                slide.classList.remove('block');
                slide.classList.remove('animate-fade-in');
            }
        });

        // Update indicators
        if (indicators.length > 0) {
            indicators.forEach((indicator, i) => {
                if (i === currentIndex) {
                    indicator.classList.add('bg-vapula-neon', 'w-8');
                    indicator.classList.remove('bg-gray-600', 'w-2', 'hover:bg-gray-400');
                } else {
                    indicator.classList.remove('bg-vapula-neon', 'w-8');
                    indicator.classList.add('bg-gray-600', 'w-2', 'hover:bg-gray-400');
                }
            });
        }
    }

    // Event Listeners
    if (prevBtn) {
        prevBtn.addEventListener('click', () => showSlide(currentIndex - 1));
    }

    if (nextBtn) {
        nextBtn.addEventListener('click', () => showSlide(currentIndex + 1));
    }

    if (indicators.length > 0) {
        indicators.forEach((indicator, index) => {
            indicator.addEventListener('click', () => showSlide(index));
        });
    }

    // Auto play (optional)
    let interval = setInterval(() => showSlide(currentIndex + 1), 5000);

    carouselContainer.addEventListener('mouseenter', () => clearInterval(interval));
    carouselContainer.addEventListener('mouseleave', () => {
        interval = setInterval(() => showSlide(currentIndex + 1), 5000);
    });

    // Initialize
    showSlide(0);
});
