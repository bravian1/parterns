let canvas, ctx, animationId;
let rotation = 0;
let isAnimating = false;

// Initialize WebAssembly
async function initWasm() {
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(
        fetch("main.wasm"),
        go.importObject
    );
    go.run(result.instance);
    
    // Hide loading message and show canvas
    document.getElementById('loading').style.display = 'none';
    canvas.style.display = 'block';
    document.querySelector('.controls').style.display = 'block';
    
    // Initial draw
    draw();
}

// Initialize the application
function init() {
    canvas = document.getElementById('canvas');
    ctx = canvas.getContext('2d');

    // Set up event listeners
    document.getElementById('depth').addEventListener('input', function(e) {
        document.getElementById('depthValue').textContent = e.target.value;
        draw();
    });

    document.getElementById('toggleAnimation').addEventListener('click', toggleAnimation);

    // Initialize WebAssembly
    initWasm().catch(err => {
        console.error("Failed to load WASM:", err);
        document.getElementById('loading').textContent = 'Failed to load WebAssembly. Please check console for errors.';
    });
}

// Main drawing function
function draw() {
    const width = canvas.width;
    const height = canvas.height;
    const centerX = width / 2;
    const centerY = height / 2;
    const radius = Math.min(width, height) * 0.4;
    const depth = parseInt(document.getElementById('depth').value);

    // Clear canvas
    ctx.clearRect(0, 0, width, height);
    
    // Set drawing style
    ctx.strokeStyle = '#2563eb';
    ctx.lineWidth = 1;

    // Calculate triangle points using WASM function
    const points = calculateTrianglePoints(centerX, centerY, radius, rotation);
    
    // Draw recursive triangle using WASM function
    drawRecursiveTriangle(
        ctx,
        points[0], points[1],  // First point
        points[2], points[3],  // Second point
        points[4], points[5],  // Third point
        depth
    );
}

// Animation functions
function animate() {
    rotation += 0.02;
    draw();
    animationId = requestAnimationFrame(animate);
}

function toggleAnimation() {
    isAnimating = !isAnimating;
    const button = document.getElementById('toggleAnimation');
    
    if (isAnimating) {
        button.textContent = 'Stop Animation';
        animate();
    } else {
        button.textContent = 'Start Animation';
        cancelAnimationFrame(animationId);
    }
}

// Initialize when the page loads
window.addEventListener('load', init);