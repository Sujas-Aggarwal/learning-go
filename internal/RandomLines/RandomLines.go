package randomlines

import "math/rand/v2"

type Line struct {
	X int32
	Y int32
}

// we need to be able to first create a program, that can give me the random lines that i want
func CreateRandomPath() Line {
	x := rand.Int32N(2*10+1) - 10 // getting any random number from 0 to 100 cause why not.
	y := rand.Int32N(2*10+1) - 10 // we can basically map them to
	line := Line{x, y}
	return line
}

// This goes in main ---
// package main

// import (
// 	"runtime"
// 	"time"

// 	randomlines "first_project/internal/RandomLines"

// 	"github.com/go-gl/gl/v4.1-core/gl"
// 	"github.com/go-gl/glfw/v3.3/glfw"
// )

// const (
// 	width  = 640
// 	height = 480
// 	// Since OpenGL uses coordinates from -1.0 to 1.0, we scale down
// 	// the integer lengths so the lines stay within the window boundaries.
// 	scale = 0.01
// )

// func init() {
// 	runtime.LockOSThread()
// }

// func main() {
// 	if err := glfw.Init(); err != nil {
// 		panic(err)
// 	}
// 	defer glfw.Terminate()

// 	glfw.WindowHint(glfw.ContextVersionMajor, 4)
// 	glfw.WindowHint(glfw.ContextVersionMinor, 1)
// 	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
// 	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

// 	window, err := glfw.CreateWindow(width, height, "Animated Random Paths", nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer window.Destroy()

// 	window.MakeContextCurrent()

// 	if err := gl.Init(); err != nil {
// 		panic(err)
// 	}

// 	// Shaders
// 	vertexShaderSource := `#version 410
// in vec2 position;
// void main() {
//     gl_Position = vec4(position, 0.0, 1.0);
// }` + "\x00"

// 	fragmentShaderSource := `#version 410
// out vec4 frag_color;
// void main() {
//     frag_color = vec4(0.0, 0.8, 1.0, 1.0); // Cyan colored path
// }` + "\x00"

// 	program := createProgram(vertexShaderSource, fragmentShaderSource)
// 	posAttrib := uint32(gl.GetAttribLocation(program, gl.Str("position\x00")))

// 	// VAO & VBO configuration
// 	var vao, vbo uint32
// 	gl.GenVertexArrays(1, &vao)
// 	gl.BindVertexArray(vao)

// 	gl.GenBuffers(1, &vbo)
// 	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

// 	gl.LineWidth(3.0)

// 	// --- Animation and Path State Variables ---
// 	// Slice to hold all permanent vertices of completed lines
// 	var permanentVertices []float32

// 	// Starting position at the center of the screen (0, 0)
// 	currentX, currentY := float32(0.0), float32(0.0)

// 	// Fetch the first random line using your custom logic
// 	line := randomlines.CreateRandomPath()
// 	targetX := currentX + float32(line.X)*scale
// 	targetY := currentY + float32(line.Y)*scale

// 	var progress float32 = 0.0
// 	const animationSpeed float32 = 100.0 // Adjust to make line drawing faster/slower
// 	lastTime := time.Now()

// 	// 4. Main Rendering Loop
// 	for !window.ShouldClose() {
// 		now := time.Now()
// 		deltaTime := float32(now.Sub(lastTime).Seconds())
// 		lastTime = now

// 		// Progress current line animation
// 		progress += animationSpeed * deltaTime

// 		if progress >= 1.0 {
// 			progress = 1.0
// 		}

// 		// Calculate current drawing head using linear interpolation
// 		animatingX := currentX + (targetX-currentX)*progress
// 		animatingY := currentY + (targetY-currentY)*progress

// 		// Build current frame's vertex array: all old lines + the one currently drawing
// 		renderVertices := make([]float32, len(permanentVertices)+4)
// 		copy(renderVertices, permanentVertices)

// 		// Append the active animating line segment
// 		renderVertices[len(permanentVertices)] = currentX
// 		renderVertices[len(permanentVertices)+1] = currentY
// 		renderVertices[len(permanentVertices)+2] = animatingX
// 		renderVertices[len(permanentVertices)+3] = animatingY

// 		// If the line is fully drawn, save it and generate the next path
// 		if progress >= 1.0 {
// 			permanentVertices = append(permanentVertices, currentX, currentY, targetX, targetY)

// 			// New starting position is the old destination
// 			currentX = targetX
// 			currentY = targetY

// 			// Call your random path function for the next segment
// 			line = randomlines.CreateRandomPath()
// 			targetX = currentX + float32(line.X)*scale
// 			targetY = currentY + float32(line.Y)*scale

// 			progress = 0.0 // Reset progress for the new segment
// 		}

// 		// Render Pass
// 		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
// 		gl.Clear(gl.COLOR_BUFFER_BIT)

// 		gl.UseProgram(program)
// 		gl.BindVertexArray(vao)

// 		// Upload the combined vertices to the GPU buffer dynamically
// 		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
// 		gl.BufferData(gl.ARRAY_BUFFER, len(renderVertices)*4, gl.Ptr(renderVertices), gl.DYNAMIC_DRAW)

// 		gl.EnableVertexAttribArray(posAttrib)
// 		gl.VertexAttribPointer(posAttrib, 2, gl.FLOAT, false, 0, nil)

// 		// Total vertices divided by 2 components (X, Y) gives us the count parameter
// 		gl.DrawArrays(gl.LINES, 0, int32(len(renderVertices)/2))

// 		window.SwapBuffers()
// 		glfw.PollEvents()
// 	}
// }

// func compileShader(source string, shaderType uint32) uint32 {
// 	shader := gl.CreateShader(shaderType)
// 	csource := gl.Str(source)
// 	gl.ShaderSource(shader, 1, &csource, nil)
// 	gl.CompileShader(shader)
// 	return shader
// }

// func createProgram(vertexSource, fragmentSource string) uint32 {
// 	vertexShader := compileShader(vertexSource, gl.VERTEX_SHADER)
// 	fragmentShader := compileShader(fragmentSource, gl.FRAGMENT_SHADER)

// 	program := gl.CreateProgram()
// 	gl.AttachShader(program, vertexShader)
// 	gl.AttachShader(program, fragmentShader)
// 	gl.LinkProgram(program)

// 	gl.DeleteShader(vertexShader)
// 	gl.DeleteShader(fragmentShader)

// 	return program
// }
