#### 文件对话框
Swing中提供了JFileChooser, 它可以显示一个文件对话框。
> JFileChooser类并不是JDialog类的子类。需要调用showOpenDialog显示打开文件的对话框，
或者调用showSaveDialog显示保存文件的对话框。接收文件的按钮被自动地标签为Open或Save.
也可以调用showDialog方法为按钮设定标签。


建立文件对话框并且获取用户选择信息的步骤:

1. 建立一个JFileChooser对象。与JDialog类的构造器不，它不需要指定父组件。
```java
JFileChooser chooser = new JFileChooser();
```

2. 调用setCurrntDiretory方法设置当前目录:
```java
chooser.setCurrentDirectory(new File("."));
```

3. 如果有一个想要作为用户选择的默认文件名，可以使用setSelectedFile方法进行指定:
```java
chooser.setSelectedFile(new File(filename));
```

4. 如果允许用户在对话框中选择多个文件:
```java
chooser.setMultiSelectionEnabled(true);
```

5. 如果想让对话框仅显示某一种类型的文件,需要设置文件过滤器。

6. 在默认情况下，用户在文件选择器中只能选择文件。如果希望选择目录，需要调用setFileSelectionMode方法.
参数值为: JFileChooser.FILES_ONLY(默认值),JFileChooser.DIRECTORIES_ONLY, 或者JFileChooser.FILES_
AND_DIRECTORIES。

7. 调用showOpenDialog或者showSaveDialog方法显示对话框。必须为这些调用提供父组件:
```java
int result = chooser.showOpenDialog(parent);
// or
int result = chooser.showSaveDialog(parent);
```
这些调用的区别是“确认按钮“的标签不同。点击“确认按钮"将完成文件选择。也可以调用showDialog
方法，并将一个显式的文件传递给确认按钮:
```java
int result = chooser.showDialog(parent, "Select");
```
仅当用户确认、取消或者离开对话框时才返回调用。返回值可以是JFileChooser.APPROVE_OPTION、
JFileChooser.CANCEL_OPTION或者JFileChooser.ERROR_OPTION。

8. 调用getSelectedFile()或者getSelectedFiles()方法获取用户选择的一个或多个文件。
这些方法将返回一个文件对象或者一组文件对象。如果需要知道文件对象名时，可以调用getPath方法:
```java
String filename = chooser.getSelectedFile().getPath();
```
> 若想限制显示的文件，需要创建一个实现了抽象类javax.swing.filechooser.FileFilter的对象。
> 文件选择器将每个文件传递给文件过滤器，只有文件过滤器接受的文件才被最终显示出来。

其实，设计专用文件过滤器非常简单，只要实现FileFiter超类中的两个方法即可:
```java
public boolean accept(File f);
public String getDescription();
```
> 需要解决同时导入javax.io包和javax.swing.filechooser包带来的名称冲突问题。最简单的方法是
> 导入javax.swing.filechooser.FileFilter, 而不要导入javax.swing.filechooser.*;

```java
chooser.setFileFilter(new FileNameExtensionFilter("Image files", "gif", "jpg"));

// or  可以为一个文件选择器安装多个过滤器:
chooser.addChoosableFileFilter(filter1);
chooser.addChoosableFileFilter(filter2);
```
如果想禁用All files过滤器，需要调用:
```java
chooser.setAcceptAllFileFilterUsed(false);
```
如果为加载和保存不同类型的文件重用一个文件选择器:
```java
chooser.resetChoosableFilters();
```
