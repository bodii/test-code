use std::collections::HashMap;
use std::io;
use regex::Regex;

fn main() {
    /*
    使用哈希 map 和 vector，创建一个文本接口来允许用户向公司的部门中增加员
    工的名字。例如，“Add Sally to Engineering” 或 “Add Amir to Sales”。接
    着让用户获取一个部门的所有员工的列表，或者公司每个部门的所有员工按照字
    典序排列的列表。
    */

    
    help_info();

    let mut employee_info: HashMap<String, Vec<String>> = HashMap::new();

    loop {
        let input = input_event(); 
        match input.as_str() {
            "help" => help_info(),
            "quit" => return,
            "one-d" => { 
                get_all_employees_in_a_department(&mut employee_info, 0);
            }, 
            "one-d-order" => { 
                get_all_employees_in_a_department(&mut employee_info, 1);
            }, 
            "one-d-reverse" => {
                get_all_employees_in_a_department(&mut employee_info, -1);
            },
            "all-order" => {
                all_employees_order(&mut employee_info, 1);
            },
            "all-reverse" => {
                all_employees_order(&mut employee_info, -1);
            },
            "d-list" => department_list(&mut employee_info),
            other => {
                let res = handle_add(other, &mut employee_info);   
                println!("result: {}", res);
                println!("-------------------\n");
            },
        }
    }
}

/*
* 添加员工命令处理
* 如： add Amir to Sales
*
*/
fn handle_add(text: &str, employee_info: &mut HashMap<String, Vec<String>>) -> bool {
    let re = Regex::new(r"(?x)^[Aa]dd\s+
    (?P<name>[\w\s]{0, 10})\s+
    to\s+
    (?P<department>[\w\s]+)$").unwrap();
    if !re.is_match(text) {
        println!("error");
        println!("-------------------\n");
        return false;
    }
    let caps = re.captures(text).unwrap();

    if caps["name"].is_empty() {
        println!("name is empty!");
        println!("-------------------\n");
        return false;
    }

    if caps["department"].is_empty() {
        println!("department is empty!");
        println!("-------------------\n");
        return false;
    }

    println!("name: |{}|, department: |{}|", 
    &caps["name"], &caps["department"]);

    let v = employee_info.entry(caps["department"].to_string())
        .or_insert(Vec::new());

    if !v.contains(&(caps["name"].to_string())) {
        v.push(caps["name"].to_string());
    } else {
        println!("\n!!!the current user already exists.\n");
    }

    println!("after added: {:?}", employee_info);
    println!("-------------------\n");

    true
}

/*
* 获取所有部门的所有员工，并按指定顺序排序
*/
fn all_employees_order(employee_info: &mut HashMap<String, 
    Vec<String>>, order: i8) {

    let mut all_employees: Vec<String> = vec![];
    let employee_info = employee_info.clone();
    for (_, mut e) in employee_info {
        if !e.is_empty() {
            all_employees.append(&mut e);
        }
    }

    if order == -1 {
        all_employees.sort_by(|a, b| b.cmp(a));
    } else {
        all_employees.sort_by(|a, b| a.cmp(b));
    }

    println!("all employees:{:?}", all_employees);
    println!("----------");
}

/*
* 获取指定部门的所有员工
* employee_info: 所有部门的员工信息
* order： 排序, 0: 按加入时的顺序; 1: 正向排序; 2: 逆向排序 
*/
fn get_all_employees_in_a_department(employee_info: &mut HashMap<String, 
    Vec<String>>, order: i8) {
    let department = input_department(order);
    
    let all_employees = employee_info.get(&department);
    if all_employees.is_none() {
        println!("department [{}] the all employees: []", department);
        println!("-------------------\n");  
        return;
    }

    let all_employees: &Vec<String> = all_employees.unwrap();
    let mut new_all_employees = all_employees.clone();
    if order == -1 {
        new_all_employees.sort_by(|a, b| b.cmp(a));
    } else if order == 1 {
        new_all_employees.sort_by(|a, b| a.cmp(b));
    }
    println!("department [{}] the all employees:{:?}",
        department, new_all_employees
    );
    println!("-------------------\n");
}

/*
* 获取公司的所有部门
*/
fn department_list(employee_info: &mut HashMap<String, Vec<String>>) {

    let employee_info = employee_info.clone();
    println!("all department: {:?}", employee_info.keys());
    println!("-------------------\n");
}

/*
* 获取命令行获取输入的部门名称
* order: 预设接下来的排序规则
*/
fn input_department(order: i8) -> String {
    let order_str: &str;
    if order == -1 {
        order_str = "逆向排序";
    } else if order == 1 {
        order_str = "正向排序";
    } else {
        order_str = "插入顺序";
    }
    println!("\n-------------------");
    println!("将获取指定部门按{order_str}的所有员工.");
    let mut department_val = String::new();

    println!("请输入部门:");
    io::stdin().read_line(&mut department_val)
        .expect("信息输入有误，请重新输入:");

    department_val.trim().to_string()
}

/*
* 获取命令面板的命令
*/
fn input_event() -> String {
    let mut input_val = String::new();

    println!("请输入你的指令（输入help显示帮助信息）:");
    io::stdin().read_line(&mut input_val).expect("信息输入有误，请重新输入:");

    input_val.trim().to_string()
}

/*
* 获取命令面板相关命令的帮助信息
*/
fn help_info() {
    println!("\n");
    println!("公司员工信息记录");
    println!("\t添加员工到指定职位");
    println!("\t\t你可以输入: Add name to department 例如：");
    println!("\t\tA[a]dd Sally to Engineering: 将Sally添加到Engineering职位下");
    println!("\t\tA[a]dd Amir to Sales: 将Amir添加到Sales职位下");
    println!("\tall-order:获取公司所有员工信息-升序"); 
    println!("\tall-reverse: 获取公司所有员工信息-降序"); 
    println!("\tone-d: 获取公司内一个部门所有员工信息-加入顺序");
    println!("\tone-d-order: 获取公司内一个部门所有员工信息-升序");
    println!("\tone-d-reverse: 获取公司内一个部门所有员工信息-降序");
    println!("\td-list: 获取公司所有的部门");
    println!("\tquit: 指令退出当前程序");
    println!("\thelp: 指令可以获取帮助信息");
    println!("\n");
}
