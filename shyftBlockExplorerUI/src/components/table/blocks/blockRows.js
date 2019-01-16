import React, { Component } from 'react';
import BlockTable from './blockTable';
import ErrorMessage from './errorMessage';
import Pagination from '../pagination/pagination';
import classes from './table.css';
import axios from "axios/index";

import {API_URL} from "../../../constants/apiURL";

class BlocksTable extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: [],
            emptyDataSet: true
        };
    }

    async componentDidMount() {
        let pageLimit= 25;
        let offset = 0;
        try {
            const response = await axios.get(`${API_URL}/get_all_blocks_length`);
            await this.setState({totalRecords: response.data});
            try {
                const response = await axios.get(`${API_URL}/get_all_blocks/${pageLimit}/${offset}`);
                if(response.data === "\n") {
                    this.setState({emptyDataSet: true})
                } else {
                    this.setState({emptyDataSet: false})
                }
                await this.setState({data: response.data});
            } catch (err) {
                console.log(err);
            }
        } catch (err) {
            console.log(err);
        }
    }

    onPageChanged = async(data) => {
        const { currentPage, totalPages, pageLimit } = data;

        const offset = (currentPage - 1) * pageLimit;

        try {
            const response = await axios.get(`${API_URL}/get_all_blocks/${pageLimit}/${offset}`);
            if(response.data === "\n") {
                this.setState({emptyDataSet: true})
            } else {
                this.setState({emptyDataSet: false})
            }
            await this.setState({data: response.data});
        } catch (err) {
            console.log(err);
        }
    };

    render() {
        let table;
        
        if(this.state.emptyDataSet === false && this.state.data.length > 0  ) {
            table = this.state.data.map((data, i) => {
                const conversion = data.Rewards / 10000000000000000000;
                return <BlockTable
                    key={`${data.TxHash}${i}`}
                    Hash={data.Hash}
                    Number={data.Number}
                    Coinbase={data.Coinbase}
                    AgeGet={data.Age}
                    GasUsed={data.GasUsed}
                    GasLimit={data.GasLimit}
                    UncleCount={data.UncleCount}
                    TxCount={data.TxCount}
                    Reward={conversion}
                    detailBlockHandler={this.props.detailBlockHandler}
                    getBlocksMined={this.props.getBlocksMined}
                />
            });
        }

        let combinedClasses = ['responsive-table', classes.table];
        return (
            <div>
                {
                     this.state.emptyDataSet === false && this.state.data.length > 0  ?
                            <table className={combinedClasses.join(' ')}>
                                <thead>
                                    <tr>
                                        <th scope="col" className={classes.thItem}> Height </th>
                                        <th scope="col" className={classes.thItem}> Block Hash </th>
                                        <th scope="col" className={classes.thItem}> Age </th>
                                        <th scope="col" className={classes.thItem}> Txn </th>
                                        <th scope="col" className={classes.thItem}> Uncles </th>
                                        <th scope="col" className={classes.thItem}> Coinbase </th>
                                        <th scope="col" className={classes.thItem}> GasUsed </th>
                                        <th scope="col" className={classes.thItem}> GasLimit </th>
                                        <th scope="col" className={classes.thItem}> Avg.GasPrice </th>
                                        <th scope="col" className={classes.thItem}> Reward </th>
                                    </tr>
                                </thead>
                                {table}
                                <div id={classes.pages}>
                                    <Pagination totalRecords={this.state.totalRecords} pageLimit={25} pageNeighbours={1} onPageChanged={this.onPageChanged} />
                                </div>
                            </table>
                    : <ErrorMessage />
                }
            </div>
        );
    }
}
export default BlocksTable;
