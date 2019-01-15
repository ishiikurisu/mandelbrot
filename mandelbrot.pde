final float FACTOR = 2 / (1 + sqrt(5));
final int MAX_ITR = 1024;
final int BLACK = 0;
final int WHITE = 255;
float FRAME_INIT_X;
float FRAME_INIT_Y;
float FRAME_END_X;
float FRAME_END_Y;

float pit(float x, float y) {
  return sqrt(x*x + y*y);
}

void firstSetting() {
  if (height > width) {
    FRAME_INIT_X = -1.0;
    FRAME_INIT_Y = -((float)height/((float)width));
    FRAME_END_X = 1.0;
    FRAME_END_Y = (float)height/((float)width);
  }
  else {
    FRAME_INIT_X = -((float)width/((float)height));
    FRAME_INIT_Y = -1.0;
    FRAME_END_X = ((float)width)/((float)height);
    FRAME_END_Y = 1.0;  
  }
}

boolean escapes(float cx, float cy) {
  boolean outlet;
  float zx = 0.0;
  float zy = 0.0;
  float temp;

  for (int i = 0; i < MAX_ITR && pit(zx, zy) <= 2; ++i) {
    temp = zx*zx - zy*zy + cx; 
    zy = 2 * zx * zy + cy;
    zx = temp;
  }

  if (pit(zx, zy) > 2) 
    outlet = true;
  else
    outlet = false;

  return outlet;
}

void zoomAt(int posX, int posY) {
  float centerX = FRAME_INIT_X + (posX * (FRAME_END_X - FRAME_INIT_X)) / (width);
  float centerY = FRAME_END_Y - (posY * (FRAME_END_Y - FRAME_INIT_Y)) / (height);
  float distX = FACTOR * (FRAME_END_X - FRAME_INIT_X) / 2;
  float distY = FACTOR * (FRAME_END_Y - FRAME_INIT_Y) / 2;
  
  FRAME_INIT_X = centerX - distX;
  FRAME_INIT_Y = centerY - distY;
  FRAME_END_X = centerX + distX;
  FRAME_END_Y = centerY + distY;
}

void mandelbrot() {
  float dy = (FRAME_END_Y - FRAME_INIT_Y) / ((float) height);
  float dx = (FRAME_END_X - FRAME_INIT_X) / ((float) width);

  for (float y = FRAME_INIT_Y; y < FRAME_END_Y; y += dy)
    for (float x = FRAME_INIT_X; x < FRAME_END_X; x += dx)
      if (!escapes(x, y))
        setPixel(x, y, true);
      else
        setPixel(x, y, false);
}

void setup() {
  size(800, 600);
  background(WHITE);
  // noLoop();
  // noStroke();
  
  firstSetting();
  int past = millis();
  print("drawing... ");
  mandelbrot();
  past = millis() - past;
  println(past + "ms");
}

void draw() {
  
}

void setPixel(float px, float py, boolean state) {
  float posX = ((px - FRAME_INIT_X) * width) / (FRAME_END_X - FRAME_INIT_X);
  float posY = ((FRAME_END_Y - py) * height) / (FRAME_END_Y - FRAME_INIT_Y);

  if (state) {
    fill(BLACK);
    stroke(BLACK);
  }
  else {
    fill(WHITE);
    stroke(WHITE);
  }

  point(posX, posY);
}

void mouseClicked() {
  zoomAt(mouseX, mouseY);
  print("zooming... ");
  mandelbrot();
  println("done");
}

void keyPressed() {
  switch (key) {
    case 'p':
    case 'P':
      float posX = mouseX;
      float posY = mouseY;
      float px = FRAME_INIT_X + (posX * (FRAME_END_X - FRAME_INIT_X)) / (width);
      float py = FRAME_END_Y - (posY * (FRAME_END_Y - FRAME_INIT_Y)) / (height); 
  
      println(mouseX + ", " + mouseY + " = " + px + " + j" + py);
    break;
    
    case 'r':
    case 'R':
      firstSetting();
      print("drawing... ");
      mandelbrot();
      println("done");
    break;
  }
}
